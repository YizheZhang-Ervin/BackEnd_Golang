package main

import (
	"fmt"
	"strconv"
)

func main() {
	storageChan := make(chan Product, 1000)
	shopChan := make(chan Product, 1000)
	exitChan := make(chan bool, 1000)
	for i := 1; i < 999; i++ {
		go Producer(storageChan, 1000)
	}
	go Logistics(storageChan, shopChan)
	go Consumer(shopChan, 10, exitChan)
	if <-exitChan {
		return
	}
}

type Product struct {
	Name string
}

func Producer(storageChan chan<- Product, count int) {
	for {
		producer := Product{"good: " + strconv.Itoa(count)}
		storageChan <- producer
		count--
		fmt.Println("produce: ", producer)
		if count < 1 {
			return
		}
	}
}

func Logistics(storageChan <-chan Product, shopChan chan<- Product) {
	for {
		product := <-storageChan
		shopChan <- product
		fmt.Println("transport: ", product)
	}
}

func Consumer(shopChan <-chan Product, count int, exitChan chan<- bool) {
	for {
		product := <-shopChan
		fmt.Println("consume: ", product)
		count--
		if count < 1 {
			exitChan <- true
			return
		}
	}
}
