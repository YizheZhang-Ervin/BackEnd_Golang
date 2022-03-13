package main

import (
    "context"
    dry_user "dry_base_plus/rpc/user"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-plugins/registry/consul/v2"
)

type User struct{}

func (*User) Add(ctx context.Context, request *dry_user.Request, response *dry_user.Response) error {
    response.C = request.A + request.B
    return nil
}

func main() {
    reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
    service := micro.NewService(micro.Registry(reg), micro.Name("user_service"))
    service.Init()
    dry_user.RegisterUserHandler(service.Server(), new(User))
    service.Run()
}
