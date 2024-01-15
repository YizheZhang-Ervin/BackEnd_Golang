package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"rpc-common/userclient"
	"userapi/internal/config"
	"userapi/internal/middlewares"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userclient.User
	UserMiddleware *middlewares.UserMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserMiddleware: middlewares.NewUserMiddleware(),
	}
}
