package utils

import (
	"github.com/go-micro/plugins/v2/registry/consul"
	"github.com/micro/go-micro/v2"
)

// 初始化micro
func InitMicro() micro.Service {
	// 初始化客户端
	consulReg := consul.NewRegistry()

	return micro.NewService(
		micro.Registry(consulReg),
	)
}
