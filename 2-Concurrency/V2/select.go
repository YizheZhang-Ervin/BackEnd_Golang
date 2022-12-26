package main

import (
	"fmt"
	"time"
)

var in2Chan chan int

func initChan(num int) {
	for i := 1; i <= num; i++ {
		in2Chan <- i
		// fmt.Println("in2Chan", i)
	}
}

func main() {
	start := time.Now()
	in2Chan = make(chan int, 100)
	go initChan(100)
	var primeChan chan int = make(chan int, 100)
	var exitChan chan bool = make(chan bool, 8)
	for i := 0; i < 8; i++ {
		go isPrime(in2Chan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 7; i++ {
			<-exitChan
			// fmt.Println("exitChan", i)
		}
	}()
label:
	for {
		select {
		case res := <-primeChan:
			fmt.Println("prime num", res)
		default:
			break label
		}
	}
	end := time.Since(start)
	fmt.Println(end)
}

func isPrime(in2Chan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
label:
	for {
		select {
		case num := <-in2Chan:
			flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				primeChan <- num
				// fmt.Println("primeChan", num)
			}
		default:
			break label
		}
	}
	fmt.Println("goroutine end")
	exitChan <- true
}
