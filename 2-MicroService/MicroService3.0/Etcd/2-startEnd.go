package main

import (
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
	serviceAddr := "192.168.29.1"
	servicePort := 8081

	s := util.NewService()
	errChan := make(chan error)
	go func() {
		err := s.RegService(serviceID, serviceName, serviceAddr+strconv.Itoa(servicePort))
		if err != nil {
			errChan <- err
			return
		}
		err = http.ListenAndServe(":8081", router)
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
	fmt.Println("发生异常,服务正在停止")
	log.Fatal(getErr)
}
