package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"userapi/internal/errorx"
	"userapi/zapx"

	"userapi/internal/config"
	"userapi/internal/handler"
	"userapi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.BizError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	// use zapx
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
