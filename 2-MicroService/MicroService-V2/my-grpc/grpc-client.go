package main

import (
	"context"
	"fmt"

	"my-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//1. 连接 grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return
	}
	defer grpcConn.Close()

	//2. 初始化 grpc 客户端
	grpcClient := pb.NewSayNameClient(grpcConn)

	// 创建并初始化Teacher 对象
	var teacher pb.Teacher
	teacher.Name = "xxName"
	teacher.Age = 18

	//3. 调用远程服务。
	t, err := grpcClient.SayHello(context.TODO(), &teacher)

	fmt.Println(t, err)
}
