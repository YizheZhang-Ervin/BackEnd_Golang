package main

import (
	"google.golang.org/grpc"
	"my-grpc-consul/pb"
	"context"
	"net"
	"fmt"
	"github.com/hashicorp/consul/api"
)

// 定义类
type Children struct {
}

// 绑定类方法, 实现借口
func (this *Children)SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello  " + p.Name
	return p, nil
}

func main()  {
	// 把grpc服务,注册到consul上.
	// 1. 初始化consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul 对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	// 3. 告诉consul, 即将注册的服务的配置信息
	reg := api.AgentServiceRegistration {
		ID:"bj38",
		Tags:[]string{"grcp", "consul"},
		Name:"grpc And Consul",
		Address:"127.0.0.1",
		Port:8800,
		Check:&api.AgentServiceCheck{
			CheckID:"consul grpc test",
			TCP:"127.0.0.1:8800",
			Timeout:"1s",
			Interval:"5s",
		},
	}

	// 4. 注册 grpc 服务到 consul 上
	consulClient.Agent().ServiceRegister(&reg)

//////////////////////以下为 grpc 服务远程调用//////////////////////////////

	// 1.初始化 grpc 对象,
	grpcServer := grpc.NewServer()

	// 2.注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	// 3.设置监听, 指定 IP/port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务启动... ")

	// 4. 启动服务
	grpcServer.Serve(listener)

}
