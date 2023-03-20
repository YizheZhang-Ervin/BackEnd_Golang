package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resMap = make(map[int]int, 10)
	lock   sync.Mutex
)

func asyncFunc(num int) {
	// 写锁
	lock.Lock()
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	// 假装比较慢
	time.Sleep(time.Second * 1)
	resMap[num] = res
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 1; i < 20; i++ {
		go asyncFunc(i)
	}
	// 等异步的写函数，然后能打印全=>无法预测要多久才能运行完
	time.Sleep(time.Second * 5)
	// 读锁
	lock.Lock()
	for k, v := range resMap {
		fmt.Printf("num:%v,factorial:%v\n", k, v)
	}
	lock.Unlock()
	end := time.Since(start)
	fmt.Println(end)
}
