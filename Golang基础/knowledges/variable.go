package knowledges

import (
	"fmt"
)

func foo() (int, string) {
	return 30, "egg"
}

func varDemo() {
	// 声明变量
	var x string
	x = "hello"
	// 声明多个变量
	var (
		y int
		z float32
	)
	y = 1
	z = 1.0
	fmt.Println(x, y, z)
	// 声明常量
	const pi = 3.14
	// 声明多个常量
	const (
		e = 2.71
		e2
	)
	// iota常量计数器，_跳过某值
	const (
		n1 = iota //0
		n2        //1
		_
		n4 //3
	)
	const (
		m1, m2 = iota + 1, iota + 2 //1,2
		m3, m4                      //2,3
	)
	// 初始化变量
	var a string = "apple"
	// 初始化多个变量 + 类型推导
	var b, c = "banana", 10
	// 短变量
	d := 20
	// 匿名变量
	f, _ := foo()
	// %d十进制，%b二进制,%o八进制,%x十六进制,%f浮点数
	fmt.Printf("%d %d %d", c, d, f)
	fmt.Printf("%s %s", a, b)
}
