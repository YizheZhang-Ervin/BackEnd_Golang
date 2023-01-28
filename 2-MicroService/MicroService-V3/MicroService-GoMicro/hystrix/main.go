package main

import (
	"context"
	"fmt"
	"go-micro/Services"
	"go-micro/Weblib"
	"go-micro/Wrappers"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

type logWrapper struct {
	client.Client
}

func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口") //这样每一次调用调用接口时都会
	return this.Client.Call(ctx, req, rsp)
}
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("localhost:8500"),
	)
	//下面两局代码是注册rpcserver调用客户端
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper),            //在注册时只需要传入方法名即可，底层会自动给这个方法传入client
		micro.WrapClient(Wrappers.NewProdsWrapper), //在注册时只需要传入方法名即可，底层会自动给这个方法传入client
	)
	prodService := Services.NewProdService("prodservice", myService.Client()) //生成的这个客户端绑定consul中存储的prodservice服务，只要调用了prodservice接口就会调用我们上面注册的中间件

	//其实下面这段代码的作用就是启动webserver的同事的时候把服务注册进去
	httpserver := web.NewService( //go-micro很灵性的实现了注册和反注册，我们启动后直接ctrl+c退出这个server，它会自动帮我们实现反注册
		web.Name("httpprodservice"),                   //注册进consul服务中的service名字
		web.Address(":8001"),                          //注册进consul服务中的端口,也是这里我们gin的server地址
		web.Handler(Weblib.NewGinRouter(prodService)), //web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
		web.Registry(consulReg),                       //注册到哪个服务器上的consul中
	)
	httpserver.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	httpserver.Run()
}
