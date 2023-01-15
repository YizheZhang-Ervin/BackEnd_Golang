package utils

import (
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
)

// 初始化micro
func InitMicro() micro.Service {
	// 初始化客户端
	consulReg := consul.NewRegistry()

	return micro.NewService(
		micro.Registry(consulReg),
	)
}
