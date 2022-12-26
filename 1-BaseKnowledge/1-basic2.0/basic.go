package main

// go run basic.go
/*
	go build basic.go
	./basic
*/

import (
	"fmt"
	"unsafe"
)

// 全局变量 (名称相同时，局部变量优先)
var (
	g1 int
	g2 string
)
var g3 int = 5

// 全局常量
const (
	c1 = "abc"
	c2 = len(c1)
	c3 = unsafe.Sizeof(c1)
)

func main() {
	fmt.Println("--------------------------test 1 变量赋值--------------------------")
	var v1 int = 123
	var v2, v3 string = "abc", "x=%d,y=%s"
	v4 := fmt.Sprintf(v3, v1, v2)
	fmt.Println(v4)

	fmt.Println("--------------------------test 2 空值--------------------------")
	var n1 int
	var n2 float64
	var n3 bool
	var n4 string
	fmt.Printf("%v %v %v %q\n", n1, n2, n3, n4)

	fmt.Println("--------------------------test 3 全局变量--------------------------")
	g1 = 1
	g2 = "a"
	fmt.Println(g1, g2)

	fmt.Println("--------------------------test4 交换变量值--------------------------")
	s1, s2 := 1, 2
	s2, s1 = s1, s2
	fmt.Println(s1, s2)
	_, s2 = s2, s1
	s3, _ := 3, 4
	fmt.Println(s1, s2, s3)

	fmt.Println("--------------------------test 5 常量--------------------------")
	const C0 int = 10
	fmt.Println(C0, c1, c2, c3)

	fmt.Println("--------------------------test 6 iota累加--------------------------")
	const (
		i1 = iota //0
		i2        //1
		i3        //2
		i4 = "ha" //独立值，iota += 1
		i5        //"ha"   iota += 1
		i6 = 100  //iota +=1
		i7        //100  iota +=1
		i8 = iota //7,恢复计数
		i9        //8
	)
	fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8, i9)

	fmt.Println("--------------------------test 7 运算符--------------------------")
	fmt.Println("--------------------------test 13 指针--------------------------")
	o1 := 2
	o2 := 3
	fmt.Println(o1 % o2)
	fmt.Println(o1 != o2)
	fmt.Println(o1 == o2 || o1 != o2)
	fmt.Println(o1 >> 1)
	o2 ^= 2
	fmt.Println(o2)
	var o3 *int    // 指针变量
	var o9 **int   // 指向指针的指针变量
	o3 = &o1       // 取存储地址
	o9 = &o3       // 指向指针o3地址
	if o3 != nil { // 判断是否空指针
		fmt.Printf("%x\n", &o1)  // 变量地址
		fmt.Printf("%x\n", o3)   // 指针变量的存储地址
		fmt.Printf("%d\n", *o3)  // 指针访问值
		fmt.Printf("%d\n", **o9) // 指向指针的指针变量
	}
	o4, o5, o6, o7 := 2, 2, 2, 2
	o8 := o4 + (o5*o6)/o7
	fmt.Println(o8)
	// 指针传参
	swap(&o1, &o2)
	fmt.Println(o1, o2)

	fmt.Println("--------------------------test 11 变量作用域--------------------------")
	g3 := 6
	fmt.Println(g3)

	fmt.Println("--------------------------test 19 类型转换--------------------------")
	var t1 int = 17
	var t2 int = 5
	var t3 float32
	// int转float32
	t3 = float32(t1) / float32(t2)
	fmt.Printf("t3的值为: %f\n", t3)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
