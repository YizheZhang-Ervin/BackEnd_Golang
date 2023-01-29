package main

import (
	"context"
	"fmt"
	"goetcd/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/product/{id:\\d+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		str := "get product ByID" + vars["id"]
		writer.Write([]byte(str))
	})
	serviceID := "p1"
	serviceName := "productservice"
	serviceAddr := "localhost:"
	servicePort := 8081

	s := util.NewService()
	errChan := make(chan error)
	httpServer := &http.Server{
		Addr:    serviceAddr + strconv.Itoa(servicePort),
		Handler: router, //router实现了ServeHttp方法，所以是Handler接口类型
	}
	go func() {
		err := s.RegService(serviceID, serviceName, serviceAddr+strconv.Itoa(servicePort))
		if err != nil {
			errChan <- err
			return
		}
		err = httpServer.ListenAndServe()
		if err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig) //当接收到信号停止
	}()
	getErr := <-errChan
	err := s.UnregService(serviceID) //执行反注册，unset该服务的key
	err = httpServer.Shutdown(context.Background())
	//可以执行一些回收工作，比如关闭数据库
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(getErr)
}
