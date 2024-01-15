package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"order/internal/config"
	"rpc-common/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
