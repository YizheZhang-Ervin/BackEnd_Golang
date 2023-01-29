package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.NewLimiter(1, 5) //1表示每次放进筒内的数量，桶内的令牌数是5，最大令牌数也是5，这个筒子是自动补充的，你只要取了令牌不管你取多少个，这里都会在每次取完后自动加1个进来，因为我们设置的是1

	for {

		if r.AllowN(time.Now(), 2) { //AllowN表示取当前的时间，这里是一次取2个，如果当前不够取两个了，本次就不取，再放一个进去，然后返回false
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("too many request")
		}
		time.Sleep(time.Second)
	}

}
