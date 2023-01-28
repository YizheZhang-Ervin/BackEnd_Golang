package main

import (
	"fmt"
	"math/rand"
	"sync"
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
		//time.Sleep(time.Second * 4)
	}
	return Product{
		ID:    101,
		Title: "Golang从入门到精通",
		Price: 12,
	}, nil
}

func RecProduct() (Product, error) {
	return Product{
		ID:    999,
		Title: "推荐商品",
		Price: 120,
	}, nil

}

func main() {
	rand.Seed(time.Now().UnixNano())
	configA := hystrix.CommandConfig{ //创建一个hystrix的config
		Timeout:               3000, //command运行超过3秒就会报超时错误
		MaxConcurrentRequests: 5,    //控制最大并发数为5，如果超过5会调用我们传入的回调函数降级
	}
	hystrix.ConfigureCommand("get_prod", configA) //hystrix绑定command
	resultChan := make(chan Product, 1)

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go (func() {
			wg.Add(1)
			defer wg.Done()
			errs := hystrix.Go("get_prod", func() error { //使用hystrix来讲我们的操作封装成command,hystrix返回值是一个chan error
				p, _ := getProduct() //这里会随机延迟0-4秒
				resultChan <- p
				return nil //这里返回的error在回调中可以获取到，也就是下面的e变量
			}, func(e error) error {
				rcp, err := RecProduct() //推荐商品,如果这里的err不是nil,那么就会忘errs中写入这个err，下面的select就可以监控到
				resultChan <- rcp
				return err
			})
			select {
			case getProd := <-resultChan:
				fmt.Println(getProd)
			case err := <-errs: //使用hystrix.Go时返回值是chan error各个协程的错误都放到errs中
				fmt.Println(err, 1)
			}
		})()
	}
	wg.Wait()
}
