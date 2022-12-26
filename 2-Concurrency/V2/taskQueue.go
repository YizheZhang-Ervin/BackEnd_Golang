package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int = 0
	var count2 int = 0
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		defer waitGroup.Done()
		defer ticker.Stop()
		for {
			t := <-ticker.C
			fmt.Println("ticker", t.Format("2006-01-02 03:04:05PM"))
			count++
			if count > 2 {
				return
			}
		}
	}()

	timer := time.NewTimer(time.Second * 1)
	go func() {
		defer waitGroup.Done()
		defer timer.Stop()
		for {
			t := <-timer.C
			fmt.Println("timer", t.Format("2006-01-02 03:04:05PM"))
			timer.Reset(time.Second)
			count2++
			if count2 > 3 {
				return
			}
		}
	}()

	waitGroup.Wait()
	fmt.Println("END")
}
