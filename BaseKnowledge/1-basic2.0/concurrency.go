package main

// go run concurrency.go
/*
	go build concurrency.go
	./basic
*/

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------------------test 22 并发--------------------------")
	// 并发
	go doSth("AAA")
	doSth("BBB")
	// 通道
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
	// 带缓冲的通道
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 关闭通道
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	// 如果上面的 c 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了
	for i := range c2 {
		fmt.Println(i)
	}
}

func doSth(s string) {
	arr := [6]int{1, 2, 3, 5}
	for v := range arr {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v, s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
