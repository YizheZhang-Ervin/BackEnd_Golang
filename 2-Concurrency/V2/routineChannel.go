package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inChan chan int

func main() {
	inChan = make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(inChan)
	go readData(inChan, exitChan)
	if <-exitChan {
		return
	}
	fmt.Println("end")
}

func writeData(inChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 50; i++ {
		tempInt := rand.Intn(4) + 10
		fmt.Printf("write||value:%v,count:%v\n", tempInt, i)
		inChan <- tempInt
	}
	close(inChan)
}

func readData(inChan chan int, exitChan chan bool) {
	var count int

	for {
		val, ok := <-inChan
		count++
		if !ok {
			break
		}
		fmt.Printf("read||value:%v,count:%v\n", val, count)
	}
	exitChan <- true
	close(exitChan)
}
