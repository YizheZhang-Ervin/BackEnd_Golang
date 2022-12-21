package main

import (
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	// 生产者
	for i := 0; i < 4; i++ {
		go doWork(c)
	}

	// 消费者
	for {
		v := <-c
		println(v)
	}
}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- rand.Int()
	}
}
