package main

import (
	"flag"
	"fmt"

	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"trans/internal/config"
	"trans/internal/server"
	"trans/internal/svc"
)

var configFile = flag.String("f", "etc/trans.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTransSvcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		reflection.Register(grpcServer)
		pb.RegisterTransSvcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
