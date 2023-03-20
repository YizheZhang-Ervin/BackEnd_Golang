package main

import (
	"context"
	"fmt"
	"log"
	"micro/sidecar"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.Handle("POST", "/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"data": "test",
			})
		})
	}
	server := &http.Server{
		Addr:    ":8088",
		Handler: ginRouter,
	}
	service := sidecar.NewService("api.jtthink.com.test")
	service.AddNode("test-"+uuid.New().String(), 8088, "localhost:8088")
	handler := make(chan error)
	go func() {
		handler <- server.ListenAndServe()
	}()
	go (func() {
		server.ListenAndServe()
	})()
	go func() {
		notify := make(chan os.Signal)
		signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		handler <- fmt.Errorf("%s", <-notify)
	}()
	//注册服务
	go func() {
		err := sidecar.RegService(service)
		if err != nil {
			handler <- err
		}
	}()
	getHandler := <-handler //阻塞一旦有错误发生，err写入信道，解除阻塞，执行反注册服务
	fmt.Println(getHandler.Error())
	//反注册服务
	err := sidecar.UnRegService(service)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}

}
