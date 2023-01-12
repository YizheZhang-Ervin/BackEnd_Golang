package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"my-grpc-consul/pb"
)

func main() {
	// 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 创建consul对象 -- (可以重新指定 consul 属性: IP/Port , 也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)

	// 服务发现. 从consuL上, 获取健康的服务
	services, _, err := consulClient.Health().Service("grpc And Consul", "grcp", true, nil)

	// 简单的负载均衡.

	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	//////////////////////以下为 grpc 服务远程调用//////////////////////////////
	// 1. 链接服务
	//grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())

	// 使用 服务发现consul 上的 IP/port 来与服务建立链接
	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())

	// 2. 初始化 grpc 客户端
	grpcClient := pb.NewHelloClient(grpcConn)

	var person pb.Person
	person.Name = "Andy"
	person.Age = 18

	// 3. 调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &person)

	fmt.Println(p, err)
}
