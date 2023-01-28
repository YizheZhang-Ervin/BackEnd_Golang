package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	myhttp "github.com/micro/go-plugins/client/http"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("106.12.72.181:23791"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	getClient := myhttp.NewClient(client.Selector(mySelector), client.ContentType("application/json"))

	//1创建request
	req := getClient.NewRequest("api.jtthink.com.test", "/v1/test", map[string]string{}) //这里的request
	//2创建response
	var rsp map[string]interface{}                         //var rsp map[string]string这里这么写也可以，因为我们的返回值是{"data":"test"}，所以都对的上
	err := getClient.Call(context.Background(), req, &rsp) //将返回值映射到map中
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
