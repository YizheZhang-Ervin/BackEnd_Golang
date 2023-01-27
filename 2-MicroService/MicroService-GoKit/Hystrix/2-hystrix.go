package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func getProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 6 { //模拟api卡顿和超时效果
		time.Sleep(time.Second * 4)
	}
	return Product{
		ID:    101,
		Title: "Golang从入门到精通",
		Price: 12,
	}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	configA := hystrix.CommandConfig{ //创建一个hystrix的config
		Timeout: 3000, //command运行超过3秒就会报超时错误
	}
	hystrix.ConfigureCommand("get_prod", configA) //hystrix绑定command
	for {
		err := hystrix.Do("get_prod", func() error { //使用hystrix来讲我们的操作封装成command
			p, _ := getProduct() //这里会随机延迟0-4秒
			fmt.Println(p)
			return nil
		}, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
