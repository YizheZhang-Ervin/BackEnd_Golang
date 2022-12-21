# Concurrency Knowledge

汇编码：go tool compile -s xx.go |grep xx.go:xx
监视内存地址的读写：go run -race xx.go
汇编+监视：go tool compile -race -S xx.go
检查锁复制：go vet
debug：DPrintf()
并发测试：go-test-many

1. sync.WaitGroup
```
waitGroup等待组(信号量)，即等待所有任务完成
add增加计数，done减少计数，wait等待计数减到0
```

2. sync.Mutex & sync.RWMutex
```
Mutex互斥锁
lock加锁，unlock解锁
defer xx.unlock是函数退出时解锁

RWMutex读写互斥锁(读锁和写锁互斥，读锁可重入)
RLock加读锁，RUnlock解读锁
Lock加写锁，Unlock解写锁
```

3. Time.Sleep
```
Sleep用于在规定持续时间内停止最新的goroutine，睡眠时间<=0此方法立即返回
```

4. sync.NewCond
```
NewCond条件变量，线程间共享的变量
Broadcast唤醒，Wait进入等待 (上下文都要带锁)
```

5. Channel
```
channel通道,用于goroutine之间发送消息
类型：
    不带缓冲
    带缓冲
特性：
    必须要有收发
    收发要在2个goroutine中
    接收会持续阻塞到发送方发数据
    每次接受一个元素
接收方式：
    阻塞接收 data := <-ch
    非阻塞接收 data, ok := <-ch
    接收任意数据，忽略接收的数据 <-ch
```