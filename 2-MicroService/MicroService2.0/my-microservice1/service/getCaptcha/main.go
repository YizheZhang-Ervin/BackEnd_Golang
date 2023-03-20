package main

import (
	"my-microservice1/service/getCaptcha/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	getCaptcha "my-microservice1/service/getCaptcha/proto/getCaptcha"

	"github.com/go-micro/plugins/v2/registry/consul"
)

func main() {
	// 初始化consul
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Address("192.168.6.108:12341"), // 防止随机生成 port
		micro.Name("getCaptcha"),
		micro.Registry(consulReg), // 添加注册
		micro.Version("latest"),
	)

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
