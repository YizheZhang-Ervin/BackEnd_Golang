package main

// go run function.go
/*
	go build function.go
	./function
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--------------------------test 10 函数--------------------------")
	// 取最大值
	maxNum := max(10, 20)
	fmt.Println(maxNum)
	// 交换值
	newStr1, newStr2 := swift("a", "z")
	fmt.Println(newStr1, newStr2)
	// 交换指针
	a, b := "b", "y"
	swiftByPointer(&a, &b)
	fmt.Println(a, b)
	// 函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
	// 函数闭包
	getNextFunc := getSequence()
	fmt.Println(getNextFunc())
	fmt.Println(getNextFunc())
	// 方法
	var s Square
	s.length = 10
	fmt.Println(s.getX())

	fmt.Println("--------------------------test 18 递归--------------------------")
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

// 函数：取最大值
func max(num1, num2 int) int {
	// num1,num2即函数局部变量
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数：交换值
func swift(str1, str2 string) (string, string) {
	return str2, str1
}

// 函数：引用传递，交换地址
func swiftByPointer(str1, str2 *string) {
	var temp string
	temp = *str1
	*str1 = *str2
	*str2 = temp
}

// 函数：函数闭包 (闭包中递增 i 变量)
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Square struct {
	length int
}

// 函数: 递归
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// 方法：
func (s Square) getX() string {
	return fmt.Sprintf("%s is %d", "Get length", s.length)
}
