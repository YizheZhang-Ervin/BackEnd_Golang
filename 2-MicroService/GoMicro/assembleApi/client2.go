package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func callAPI(s selector.Selector) {
	myCli := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myCli.NewRequest("prodservice", "/v1/prods", map[string]int{"SIZE": 3}) //使用生成的pb文件中的结构体作为参数封装到请求体中

	var resp Models.ProdListResponse //这里使用生成的response对象，这样就避免了我们在传递时候参数类型的不灵活，也就解决了上节课的问题
	err := myCli.Call(context.Background(), req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin), //设置查询策略，这里是轮询
	)
	callAPI(mySelector)
}
