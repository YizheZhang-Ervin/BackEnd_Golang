package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

// Part1-1：每个(Map或者Reduce)Task有分为idle, in-progress, completed 三种状态
type MasterTaskStatus int

const (
	Idle MasterTaskStatus = iota
	InProgress
	Completed
)

// Part1-5：此外我将task和master的状态合并成一个State。task和master的状态应该一致。如果在Reduce阶段收到了迟来MapTask结果，应该直接丢弃。
type State int

const (
	Map State = iota
	Reduce
	Exit
	Wait
)

// Part1-3：此外Master存储Map任务产生的R个中间文件的信息
type Master struct {
	TaskQueue     chan *Task          // 等待执行的task
	TaskMeta      map[int]*MasterTask // 当前所有task的信息
	MasterPhase   State               // Master的阶段
	NReduce       int
	InputFiles    []string
	Intermediates [][]string // Map任务产生的R个中间文件的信息
}

// Part1-2： Master存储这些Task的信息。与论文不同的是，这里我并没有保留worked的ID，因此master不会主动向worker发送心跳检测
type MasterTask struct {
	TaskStatus    MasterTaskStatus
	StartTime     time.Time
	TaskReference *Task
}

// Part1-4：Map和Reduce的Task应该负责不同的事情，但是在实现代码的过程中发现同一个Task结构完全可以兼顾两个阶段的任务
type Task struct {
	Input         string
	TaskState     State
	NReducer      int
	TaskNumber    int
	Intermediates []string
	Output        string
}

var mu sync.Mutex

//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// Part2-9：master确认所有ReduceTask都已经完成，转入Exit阶段，终止所有master和worker goroutine
// main/mrmaster.go calls Exit() periodically to find out
// if the entire job has finished.
//
func (m *Master) Done() bool {
	mu.Lock()
	defer mu.Unlock()
	ret := m.MasterPhase == Exit
	return ret
}

// Part 2-1：启动master
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeMaster(files []string, nReduce int) *Master {
	m := Master{
		TaskQueue:     make(chan *Task, max(nReduce, len(files))),
		TaskMeta:      make(map[int]*MasterTask),
		MasterPhase:   Map,
		NReduce:       nReduce,
		InputFiles:    files,
		Intermediates: make([][]string, nReduce),
	}

	// 切成16MB-64MB的文件
	// 创建map任务
	m.createMapTask()

	// 一个程序成为master，其他成为worker
	//这里就是启动master 服务器就行了，
	//拥有master代码的就是master，别的发RPC过来的都是worker
	m.server()
	// 启动一个goroutine 检查超时的任务
	go m.catchTimeOut()
	return &m
}

// Part2-10：周期性检查task是否完成。将超时未完成的任务，交给新的worker，backup执行
func (m *Master) catchTimeOut() {
	for {
		time.Sleep(5 * time.Second)
		mu.Lock()
		if m.MasterPhase == Exit {
			mu.Unlock()
			return
		}
		for _, masterTask := range m.TaskMeta {
			if masterTask.TaskStatus == InProgress && time.Now().Sub(masterTask.StartTime) > 10*time.Second {
				m.TaskQueue <- masterTask.TaskReference
				masterTask.TaskStatus = Idle
			}
		}
		mu.Unlock()
	}
}

func (m *Master) createMapTask() {
	// 根据传入的filename，每个文件对应一个map task
	for idx, filename := range m.InputFiles {
		taskMeta := Task{
			Input:      filename,
			TaskState:  Map,
			NReducer:   m.NReduce,
			TaskNumber: idx,
		}
		m.TaskQueue <- &taskMeta
		m.TaskMeta[idx] = &MasterTask{
			TaskStatus:    Idle,
			TaskReference: &taskMeta,
		}
	}
}

func (m *Master) createReduceTask() {
	m.TaskMeta = make(map[int]*MasterTask)
	for idx, files := range m.Intermediates {
		taskMeta := Task{
			TaskState:     Reduce,
			NReducer:      m.NReduce,
			TaskNumber:    idx,
			Intermediates: files,
		}
		m.TaskQueue <- &taskMeta
		m.TaskMeta[idx] = &MasterTask{
			TaskStatus:    Idle,
			TaskReference: &taskMeta,
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Part2-2：master监听worker RPC调用，分配任务
// master等待worker调用
func (m *Master) AssignTask(args *ExampleArgs, reply *Task) error {
	// assignTask就看看自己queue里面还有没有task
	mu.Lock()
	defer mu.Unlock()
	if len(m.TaskQueue) > 0 {
		//有就发出去
		*reply = *<-m.TaskQueue
		// 记录task的启动时间
		m.TaskMeta[reply.TaskNumber].TaskStatus = InProgress
		m.TaskMeta[reply.TaskNumber].StartTime = time.Now()
	} else if m.MasterPhase == Exit {
		*reply = Task{TaskState: Exit}
	} else {
		// 没有task就让worker 等待
		*reply = Task{TaskState: Wait}
	}
	return nil
}

// Part2-7：master收到完成后的Task
func (m *Master) TaskCompleted(task *Task, reply *ExampleReply) error {
	//更新task状态
	mu.Lock()
	defer mu.Unlock()
	// Part2-11：从第一个完成的worker获取结果，将后序的backup结果丢弃
	if task.TaskState != m.MasterPhase || m.TaskMeta[task.TaskNumber].TaskStatus == Completed {
		// 因为worker写在同一个文件磁盘上，对于重复的结果要丢弃
		return nil
	}
	m.TaskMeta[task.TaskNumber].TaskStatus = Completed
	go m.processTaskResult(task)
	return nil
}

func (m *Master) processTaskResult(task *Task) {
	mu.Lock()
	defer mu.Unlock()
	switch task.TaskState {
	case Map:
		//收集intermediate信息
		for reduceTaskId, filePath := range task.Intermediates {
			m.Intermediates[reduceTaskId] = append(m.Intermediates[reduceTaskId], filePath)
		}
		if m.allTaskDone() {
			//获得所以map task后，进入reduce阶段
			m.createReduceTask()
			m.MasterPhase = Reduce
		}
	case Reduce:
		if m.allTaskDone() {
			//获得所以reduce task后，进入exit阶段
			m.MasterPhase = Exit
		}
	}
}

func (m *Master) allTaskDone() bool {
	for _, task := range m.TaskMeta {
		if task.TaskStatus != Completed {
			return false
		}
	}
	return true
}
