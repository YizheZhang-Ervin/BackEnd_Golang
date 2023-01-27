package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.NewLimiter(1, 5) //1表示每次放进筒内的数量，桶内的令牌数是5，最大令牌数也是5，这个筒子是自动补充的，你只要取了令牌不管你取多少个，这里都会在每次取完后自动加1个进来，因为我们设置的是1
	ctx := context.Background()

	for {
		err := r.WaitN(ctx, 2) //每次消耗2个，放入一个，消耗完了还会放进去，因为初始是5个，所以这段代码再执行到第4次的时候筒里面就空了，如果当前不够取两个了，本次就不取，再放一个进去，然后返回false
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format("2016-01-02 15:04:05"))
		time.Sleep(time.Second)
	}

}
