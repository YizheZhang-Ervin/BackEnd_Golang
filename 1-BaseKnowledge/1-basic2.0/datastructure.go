package main

// go run datastructure.go
/*
	go build datastructure.go
	./basic
*/

import "fmt"

func main() {
	fmt.Println("--------------------------test 12 数组--------------------------")
	// 一维数组
	var a1 [5]float32
	a1[0] = 1.0
	var a2 = [5]float32{2.0, 3.0, 4.0, 5.0, 6.0}
	a3 := [3]float32{7.0, 8.0, 9.0}
	a4 := [...]float32{10.0, 11.0}
	a5 := [3]float32{1: 12.0, 2: 13.0}
	a5[2] = 14.0
	var a6 float32 = a5[0]
	fmt.Println(a1, a2, a3, a4, a5, a6)
	// 多维数组
	a7 := [][]int{}
	a8 := []int{1, 2, 3}
	a7 = append(a7, a8)
	fmt.Println(a7[0][0])
	a9 := [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(a9[0][0])
	// 数组形参
	getAvg(a5)
	// 指针数组
	var a10 [3]*int
	var i int
	for i = 0; i < 3; i++ {
		a10[i] = &a8[i]
	}
	fmt.Println(a10)

	fmt.Println("--------------------------test 14 结构体--------------------------")
	b1 := Books{"title001", "author001", "subject001", 001}
	b2 := Books{title: "title002", author: "author002", subject: "subject002", book_id: 002}
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b1.title)
	getBook(b2)
	getBook2(&b2)

	fmt.Println("--------------------------test 15 切片--------------------------")
	var c1 []int = make([]int, 3)
	if c1 == nil {
		fmt.Println("c1 nil")
	}
	c2 := []int{1, 2, 3}
	fmt.Println(len(c2), cap(c2))
	fmt.Println(c2[:1])
	c1 = c2[:]
	fmt.Printf("%v\n", c1)
	// 切片增加容量
	c2 = append(c2, 1, 2, 3)
	fmt.Println(c2, cap(c2))
	// 创新切片
	c3 := make([]int, len(c2), (cap(c2))*2)
	copy(c3, c2)
	fmt.Println(c3, cap(c3))

	fmt.Println("--------------------------test 16 范围--------------------------")
	var d1 = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range d1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// 数组索引
	for i, num := range d1 {
		fmt.Println(i, num)
	}
	// map键值对
	d2 := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range d2 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// (字符索引，字符Unicode的值)
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println("--------------------------test 17 集合Map--------------------------")
	var e1 map[string]string
	e1 = make(map[string]string)
	e1["key1"] = "value1"
	val, ok := e1["key1"]
	if ok {
		fmt.Println(val)
	}
	// 删除
	e2 := map[string]string{"k1": "v1", "k2": "v2"}
	delete(e2, "k2")
	fmt.Println(e2)
}

func getAvg(arr [3]float32) {
	fmt.Println(arr[1])
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func getBook(book Books) {
	fmt.Println(book.title)
}

func getBook2(book *Books) {
	fmt.Printf("%s", book.author)
}
