package main

import "fmt"

func main() {
	go addMapVal("abc")
	fmt.Println("END")
}

func addMapVal(str string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PANIC!!!", err)
		}
	}()
	var testMap map[int]string
	testMap[0] = str
}
