package main

import (
	"context"
	"fmt"
	mygomicrogin "my-gomicro-gin/proto/my-gomicro-server"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v2/registry/consul"
	"github.com/micro/go-micro/v2"
)

func CallRemote(ctx *gin.Context) {

	// 初始化服务发现 consul
	consulReg := consul.NewRegistry()

	// 初始化micro服务对象, 指定consul 为服务发现
	service := micro.NewService(
		micro.Registry(consulReg),
	)

	// 1. 初始化客户端
	microClient := mygomicrogin.NewMyGomicroServerService("go.micro.srv.my-gomicro-gin", service.Client())
	fmt.Println()

	// 2. 调用远程服务
	resp, err := microClient.Call(context.TODO(), &mygomicrogin.Request{
		Name: "xxx",
	})
	if err != nil {
		fmt.Println("call err:", err)
		return
	}

	// 为了方便查看, 在打印之前将结果返回给浏览器
	ctx.Writer.WriteString(resp.Msg)

	fmt.Println(resp, err)
}

func main() {
	// 1. 初始化路由 -- 官网:初始化 web 引擎
	router := gin.Default()

	// 2. 做路由匹配
	router.GET("/", CallRemote)

	// 3. 启动运行
	router.Run(":8080")
}
