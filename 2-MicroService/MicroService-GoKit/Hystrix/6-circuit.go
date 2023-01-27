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
		time.Sleep(time.Second * 10)
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
		Timeout:                3000,                  //command运行超过3秒就会报超时错误，并且在一个统计窗口内处理的请求数量达到阈值会调用我们传入的降级回调函数
		MaxConcurrentRequests:  5,                     //控制最大并发数为5，并且在一个统计窗口内处理的请求数量达到阈值会调用我们传入的降级回调函数
		RequestVolumeThreshold: 5,                     //判断熔断的最少请求数，默认是5；只有在一个统计窗口内处理的请求数量达到这个阈值，才会进行熔断与否的判断
		ErrorPercentThreshold:  5,                     //判断熔断的阈值，默认值5，表示在一个统计窗口内有50%的请求处理失败，比如有20个请求有10个以上失败了会触发熔断器短路直接熔断服务
		SleepWindow:            int(time.Second * 10), //熔断器短路多久以后开始尝试是否恢复，这里设置的是10
	}
	hystrix.ConfigureCommand("get_prod", configA) //hystrix绑定command
	c, _, _ := hystrix.GetCircuit("get_prod")     //返回值有三个，第一个是熔断器指针,第二个是bool表示是否能够取到，第三个是error
	resultChan := make(chan Product, 1)

	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			errs := hystrix.Do("get_prod", func() error { //使用hystrix来讲我们的操作封装成command,hystrix返回值是一个chan error
				p, _ := getProduct() //这里会随机延迟0-4秒
				fmt.Println("hello")
				resultChan <- p
				return nil //这里返回的error在回调中可以获取到，也就是下面的e变量
			}, func(e error) error {
				fmt.Println("hello")
				rcp, err := RecProduct() //推荐商品,如果这里的err不是nil,那么就会忘errs中写入这个err，下面的select就可以监控到
				resultChan <- rcp
				return err
			})
			if errs != nil { //这里errs是error接口，但是使用hystrix.Go异步执行时返回值是chan error各个协程的错误都放到errs中
				fmt.Println(errs)
			} else {
				select {
				case prod := <-resultChan:
					fmt.Println(prod)
				}
			}

			fmt.Println(c.IsOpen())       //判断熔断器是否打开，一旦打开所有的请求都会走fallback
			fmt.Println(c.AllowRequest()) //判断是否允许请求服务，一旦打开
		}()
	}
	wg.Wait()
}
