# MapReduce

## MapReduce的执行流程
1. 启动MapReduce, 将输入文件切分成大小在16-64MB之间的文件。然后在一组多个机器上启动用户程序
2. 其中一个副本将成为master, 余下成为worker. master给worker指定任务（M个map任务，R个reduce任务）。master选择空闲的worker给予map或reduce任务
3. Map worker 接收切分后的input，执行Map函数，将结果缓存到内存
4. 缓存后的中间结果会周期性的写到本地磁盘，并切分成R份（reducer数量）。R个文件的位置会发送给master, master转发给reducer
5. Reduce worker 收到中间文件的位置信息，通过RPC读取。读取完先根据中间<k, v>排序，然后按照key分组、合并。
6. Reduce worker在排序后的数据上迭代，将中间<k, v> 交给reduce 函数处理。最终结果写给对应的output文件（分片）
7. 所有map和reduce任务结束后，master唤醒用户程序

## MapReduce的实现
1. 每个(Map或者Reduce)Task有分为idle, in-progress, completed 三种状态。
2. Master存储这些Task的信息。与论文不同的是，这里我并没有保留worked的ID，因此master不会主动向worker发送`心跳检测`
3. 此外Master存储Map任务产生的R个中间文件的信息。
4. Map和Reduce的Task应该负责不同的事情，但是在实现代码的过程中发现同一个Task结构完全可以兼顾两个阶段的任务。
5. 此外我将task和master的状态合并成一个State。task和master的状态应该一致。如果在Reduce阶段收到了迟来MapTask结果，应该直接丢弃。

## MapReduce执行过程的实现
1. 启动master
2. master监听worker RPC调用，分配任务
3. 启动worker
4. worker向master发送RPC请求任务
5. worker获得MapTask，交给mapper处理
6. worker任务完成后通知master
7. master收到完成后的Task
8. 转入Reduce阶段，worker获得ReduceTask，交给reducer处理
9. master确认所有ReduceTask都已经完成，转入Exit阶段，终止所有master和worker goroutine
10. 上锁
```
master跟多个worker通信，master的数据是共享的，其中TaskMeta, Phase, Intermediates, TaskQueue都有读写发生。TaskQueue使用channel实现，自己带锁。只有涉及Intermediates, TaskMeta, Phase的操作需要上锁

另外go -race并不能检测出所有的datarace。我曾一度任务Intermediates写操作发生在map阶段，读操作发生在reduce阶段读，逻辑上存在barrier，所以不会有datarace. 但是后来想到两个write也可能造成datarace，然而Go Race Detector并没有检测出来。
```
11. carsh处理
```
test当中有容错的要求，不过只针对worker
1. 周期性向worker发送心跳检测
	如果worker失联一段时间，master将worker标记成failed
	worker失效之后
		已完成的map task被重新标记为idle
		已完成的reduce task不需要改变
		原因是：map的结果被写在local disk，worker machine 宕机会导致map的结果丢失；reduce结果存储在GFS，不会随着machine down丢失
2. 对于in-progress 且超时的任务，启动backup执行
3. 周期性检查task是否完成。将超时未完成的任务，交给新的worker，backup执行
4. 从第一个完成的worker获取结果，将后序的backup结果丢弃
```




## 测试
- 开启race detector，执行test-mr.sh

```bash
$ ./test-mr.sh        
*** Starting wc test.
--- wc test: PASS
*** Starting indexer test.
--- indexer test: PASS
*** Starting map parallelism test.
--- map parallelism test: PASS
*** Starting reduce parallelism test.
--- reduce parallelism test: PASS
*** Starting crash test.
--- crash test: PASS
*** PASSED ALL TESTS
```
