package main

import (
	"fmt"
)

var intChan chan int

func main() {
	// 整数类型通道
	intChan = make(chan int, 10)
	intChan <- 1
	fmt.Printf("value:%v,address:%v,size:%v,volume:%v\n", <-intChan, &intChan, len(intChan), cap(intChan))

	// 接口空类型通道
	allChan := make(chan interface{}, 10)
	allChan <- dog{Name: "x", Color: "yellow"}
	allChan <- "y"
	obj1 := <-allChan
	// 类型断言
	dog1 := obj1.(dog)
	fmt.Printf("%v%v", dog1.Name, dog1.Color)
	fmt.Printf("%v", <-allChan)

	allChan <- dog{Name: "x2", Color: "blue"}
	allChan <- "y2"
	// 循环取值需要先close通道,close之后不能再加入值
	close(allChan)
	// 循环1
	// for val := range allChan {
	// 	fmt.Println(val)
	// }

	// 循环2
	for {
		val, ok := <-allChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
}

type dog struct {
	Name  string
	Color string
}
