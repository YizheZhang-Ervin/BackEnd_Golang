package main

import (
	"google.golang.org/grpc"
	"my-grpc/pb"
	"context" // 上下文. --- goroutine (go程) 之间用来进行数据传递 API 包
	"net"
	"fmt"
)

// 定义类
type Children struct {
}

// 按接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += " is Sleeping"
	return t, nil
}

func main() {
	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()

	//2. 注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))

	//3. 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	//4. 启动服务。---- serve()
	grpcServer.Serve(listener)
}
