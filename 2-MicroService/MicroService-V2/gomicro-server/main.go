package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"test66/handler"
	"test66/subscriber"

	test66 "test66/proto/test66"
	"github.com/micro/go-micro/registry/consul"
)

func main() {

	// 初始化服务发现
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test66"),
		micro.Registry(consulReg),
		micro.Version("latest"),
	)

	// Initialise service
	//service.Init()

	// Register Handler
	test66.RegisterTest66Handler(service.Server(), new(handler.Test66))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test66", service.Server(), new(subscriber.Test66))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test66", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
