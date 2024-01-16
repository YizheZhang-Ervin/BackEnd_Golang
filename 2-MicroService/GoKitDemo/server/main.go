package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"server/TES"
)

func main() {
	var (
		httpAddr = flag.String("httpAddr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	// 日志
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// 服务 & 中间件
	var s TES.Service
	{
		s = TES.NewInmemService()
		s = TES.LoggingMiddleware(logger)(s)
	}

	// 路由
	var h http.Handler
	{
		h = TES.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	}

	// 中断
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// 启动
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	// 退出日志
	logger.Log("exit", <-errs)
}
