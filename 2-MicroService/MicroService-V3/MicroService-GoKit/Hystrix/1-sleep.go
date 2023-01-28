package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func getProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 6 { //模拟api卡顿和超时效果
		time.Sleep(time.Second * 3)
	}
	return Product{
		ID:    101,
		Title: "Golang从入门到精通",
		Price: 12,
	}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		p, _ := getProduct()
		fmt.Println(p)
		time.Sleep(time.Second)
	}
}
