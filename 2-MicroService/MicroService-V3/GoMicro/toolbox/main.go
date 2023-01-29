// micro --registry=etcd --registry_address=106.12.72.181:23791 list services
// micro --registry=etcd --registry_address=106.12.72.181:23791 get service test.xiahualou.com
// micro --registry=etcd --registry_address=106.12.72.181:23791 call test.xiahualou.com TestService.Call "{\"id\":3}"
// 调用的时候必须要加上Endpoint，传入的json参数key要用双引号括起来，反引号转义
// micro --registry=etcd --registry_address=106.12.72.181:23791 web

package main

import (
	"micro/Services"
	"micro/ServicesImpl"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	//consulReg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	etcdReg := etcd.NewRegistry(registry.Addrs("106.12.72.181:23791")) //注册服务到etcd中
	myservice := micro.NewService(
		micro.Name("test.xiahualou"+".com"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
	)
	Services.RegisterTestServiceHandler(myservice.Server(), new(ServicesImpl.TestService))
	myservice.Run()
}
