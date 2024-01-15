package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	// 从etcd获取userrpc地址
	UserRpc zrpc.RpcClientConf
}
