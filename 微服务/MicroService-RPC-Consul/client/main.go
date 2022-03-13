package main

import (
    "context"
    dry_user "dry_base_plus/rpc/user"
    "fmt"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
    reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
    service := micro.NewService(micro.Registry(reg), micro.Name("client_service"))
    service.Init()
    userService := dry_user.NewUserService("user_service", service.Client())
    response, _ := userService.Add(context.TODO(), &dry_user.Request{A: 100, B: 900})
    fmt.Println(response.C)
}
