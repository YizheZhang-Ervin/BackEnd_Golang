# 并发编程

## 子程序
0. 检查cpu核数 main.go
1. 互斥锁 lock.go
2. 通道 channel.go
3. 协程+通道 routineChannel.go
4. 通道监听(label/select/case) select.go
5. 异常捕捉(defer/panic) deferPanic.go
6. 生产者消费者 producerConsumer.go
7. 定时器(timer) timer.go
8. 定时任务(ticker) scheduleTask.go
9. 任务队列(timer/ticker) taskQueue.go

## 使用
```
# 模块
go mod init xx

# 依赖
go mod tidy

# 单个运行
go run xx.go 或 go build xx.go + ./xx.exe

# 整体运行
go build . 然后 ./xx.exe

# 检查资源竞争
go build -race xx.go
```