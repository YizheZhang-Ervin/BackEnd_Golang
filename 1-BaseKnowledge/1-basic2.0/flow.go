package main

// go run flow.go
/*
	go build flow.go
	./flow
*/

import "fmt"

func main() {
	fmt.Println("--------------------------test 8 条件语句--------------------------")
	// if-else
	a1 := 9
	if a1 >= 1 && a1 <= 5 {
		fmt.Println("1<=a1<=5")
	} else if a1 > 5 {
		fmt.Println("a1>5")
	} else {
		fmt.Println("a1<1")
	}
	// switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T\n", i)
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	default:
		fmt.Printf("未知型\n")
	}

	// switch-fallthrough
	switch {
	case true:
		fmt.Println("switch Line 1")
		fallthrough // 用于强制执行下一个case
	case false:
		fmt.Println("switch Line 2")
	}

	// select-channel 【多个 case 都可以运行，Select 会随机公平地选出一个执行】
	var a2, a3, a4 chan int
	var a5, a6 int
	select {
	case a5 = <-a2:
		fmt.Println("received ", a5)
	case a3 <- a6:
		fmt.Println("sent ", a6)
	case a7, ok := (<-a4):
		if ok {
			fmt.Println("received ", a7)
		} else {
			fmt.Println("closed")
		}
	default:
		fmt.Println("no communication")
	}

	fmt.Println("--------------------------test 9 循环语句--------------------------")
	// for
	b1 := 0
	for i := 0; i <= 10; i++ {
		b1 += i
	}
	fmt.Println(b1)

	// for condition - break
	b2 := 10
	for b2 > 5 {
		b2++
		if b2 > 15 {
			break
		}
	}

	// for - break with label (打断多重循环)
re:
	for {
		for {
			b2++
			break re
		}
	}

	// for range (K,V)
	b3 := [6]int{1, 2, 3, 5}
	for i, x := range b3 {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// for range (K)
	b4 := make(map[int]float32)
	b4[1] = 1.0
	b4[2] = 2.0
	b4[3] = 3.0
	b4[4] = 4.0
	for key := range b4 {
		fmt.Printf("key is: %d\n", key)
	}

	// for range (V)
	for _, value := range b4 {
		fmt.Printf("value is: %f\n", value)
	}

	// goto 跳到指定代码处
	var b5 int = 10
LOOP:
	for b5 < 20 {
		if b5 == 15 {
			b5 = b5 + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", b5)
		b5++
	}
}
