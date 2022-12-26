package main

import (
	"fmt"
	"math/rand"
	"time"
)

var flag bool = isStopTimer()

func main() {
	timer := time.NewTimer(time.Second * 3)

	if flag {
		timer.Stop()
	} else {
		t := <-timer.C
		fmt.Println(t)
	}

	// t := <-time.After(time.Second * 3)
	// fmt.Println(t)
}

func isStopTimer() bool {
	rand.Seed(time.Now().UnixNano())
	tempInt := rand.Intn(2) + 16
	if tempInt >= 18 {
		fmt.Println("FIND")
		return true
	} else {
		return false
	}
}
