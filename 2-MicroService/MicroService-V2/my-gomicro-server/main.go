package main

import (
	"my-gomicro-server/handler"
	"my-gomicro-server/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	mygomicroserver "my-gomicro-server/proto/my-gomicro-server"
	// "github.com/go-micro/plugins/v2/registry/consul"
	// "github.com/go-micro/plugins/v2/registry"
)

func main() {
	// 默认服务发现是mdns(如果不用consul)
	// 初始化服务发现
	// consulReg := consul.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{
	// 		"192.168.6.108:8800",
	// 	}
	// })

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.my-gomicro-server"),
		micro.Version("latest"),
		// micro.Registry(consulReg),
	)

	// Initialise service,优先级比NewService高,代码运行期使用
	service.Init()

	// Register Handler 注册服务
	mygomicroserver.RegisterMyGomicroServerHandler(service.Server(), new(handler.MyGomicroServer))

	// Register Struct as Subscriber 发布订阅
	micro.RegisterSubscriber("go.micro.service.my-gomicro-server", service.Server(), new(subscriber.MyGomicroServer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
