package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 最大cpu数
	runtime.GOMAXPROCS(16)
	// 当前cpu数
	fmt.Println("MAX CPU NUM:", runtime.NumCPU())
}
