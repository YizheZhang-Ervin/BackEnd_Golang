package main

import (
	"TES-client/Services"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	consulapi "github.com/hashicorp/consul/api"
)

// 无consul的main
func main() {
	tgt, _ := url.Parse("http://127.0.0.1:8080")
	//创建一个直连client，这里我们必须写两个func,一个是如何请求,一个是响应我们怎么处理
	client := httptransport.NewClient("GET", tgt, Services.GetUserInfo_Request, Services.GetUserInfo_Response)
	/*
	   func GenUserEnPoint(userService IUserService) endpoint.Endpoint {
	       return func(ctx context.Context, request interface{}) (response interface{}, err error) {
	           r := request.(UserRequest) //通过类型断言获取请求结构体
	           result := "nothings"
	           if r.Method == "GET" {
	               result = userService.GetName(r.Uid)
	           } else if r.Method == "DELETE" {
	               err := userService.DelUser(r.Uid)
	               if err != nil {
	                   result = err.Error()
	               } else {
	                   result = fmt.Sprintf("userid为%d的用户已删除", r.Uid)
	               }
	           }
	           return UserResponse{Result: result}, nil
	       }
	   }
	*/
	getUserInfo := client.Endpoint() //通过这个拿到了定义在服务端的endpoint也就是上面这段代码return出来的函数，直接在本地就可以调用服务端的代码

	ctx := context.Background() //创建一个上下文

	//执行
	res, err := getUserInfo(ctx, Services.UserRequest{Uid: 101}) //使用go-kit插件来直接调用服务
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := res.(UserResponse)
	fmt.Println(userinfo.Result)

}

// 有consul的main
func main2() {
	//第一步创建client
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"
	api_client, _ := consulapi.NewClient(config)
	client := consul.NewClient(api_client)

	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stdout)
	var Tag = []string{"primary"}
	//第二步创建一个consul的实例
	//最后的true表示只有通过健康检查的服务才能被得到
	instancer := consul.NewInstancer(client, logger, "userservice", Tag, true)
	//factory定义了如何获得服务端的endpoint,这里的service_url是从consul中读取到的service的address我这里是192.168.3.14:8000
	factory := func(service_url string) (endpoint.Endpoint, io.Closer, error) {
		//server ip +8080真实服务的地址
		tart, _ := url.Parse("http://" + service_url)
		//GetUserInfo_Request里面定义了访问哪一个api把url拼接成了http://192.168.3.14:8000/v1/user/{uid}的形式
		return httptransport.NewClient("GET", tart, Services.GetUserInfo_Request, Services.GetUserInfo_Response).Endpoint(), nil, nil

	}
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	endpoints, _ := endpointer.Endpoints()
	fmt.Println("服务有", len(endpoints), "条")

	// 非负载均衡
	// fmt.Println("服务有", len(endpoints), "条")
	// getUserInfo := endpoints[0] //写死获取第一个
	// ctx := context.Background() //第三步：创建一个context上下文对象

	// //第四步：执行
	// res, err := getUserInfo(ctx, Services.UserRequest{Uid: 101})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// //第五步：断言，得到响应值
	// userinfo := res.(Services.UserResponse)
	// fmt.Println(userinfo.Result)

	// 负载均衡
	// 随机
	// mylb := lb.NewRandom(endpointer, time.Now().UnixNano()) //使用go-kit自带的轮询
	// 轮询
	mylb := lb.NewRoundRobin(endpointer) //使用go-kit自带的轮询
	for {
		getUserInfo, err := mylb.Endpoint() //写死获取第一个
		ctx := context.Background()         //第三步：创建一个context上下文对象

		//第四步：执行
		res, err := getUserInfo(ctx, Services.UserRequest{Uid: 101})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//第五步：断言，得到响应值
		userinfo := res.(Services.UserResponse)
		fmt.Println(userinfo.Result)
	}
}

// 用熔断器的main
func main3() {
	configA := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  5,
		RequestVolumeThreshold: 3,
		SleepWindow:            int(time.Second * 10),
		ErrorPercentThreshold:  20,
	}

	hystrix.ConfigureCommand("getuser", configA)
	err := hystrix.Do("getuser", func() error {
		res, err := util.GetUser() //调用方法
		fmt.Println(res)
		return err
	}, func(e error) error {
		fmt.Println("降级用户")
		return e
	})
	if err != nil {
		log.Fatal(err)
	}
}
