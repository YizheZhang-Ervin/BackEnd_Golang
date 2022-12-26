package main

import (
	"fmt"
	"time"
)

func main() {
	var count int = 0

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			t := <-ticker.C
			fmt.Println(t.Format("2006-01-02 03:04:05PM"))
			count++
			if count > 2 {
				ticker.Stop()
			}
		}
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("END")
}
