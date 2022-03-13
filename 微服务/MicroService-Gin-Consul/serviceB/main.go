package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	consul "github.com/micro/go-plugins/registry/consul/v2"
)

var reg registry.Registry

func init() {
	reg = consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
}

func InitWeb() *gin.Engine {
	r := gin.Default()
	r.GET("/other/test", func(c *gin.Context) {
		content := Call(c.Query("x"), c.Query("y"))
		c.String(200, content)
	})
	return r
}

func Call(x, y string) string {
	address := GetServiceAddress("serviceA")
	url := fmt.Sprintf("http://"+address+"/api/add?x=%s&y=%s", x, y)
	response, _ := http.Get(url)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	return string(content)
}

func GetServiceAddress(name string) (address string) {
	list, _ := reg.GetService(name)
	var services []*registry.Service
	for _, value := range list {
		services = append(services, value)
	}
	next := selector.RoundRobin(services)
	if node, err := next(); err == nil {
		address = node.Address
	}
	return
}

func main() {
	service := web.NewService(
		web.Name("serviceB"),
		web.Address(":50001"),
		web.Handler(InitWeb()),
		web.Registry(reg),
	)
	_ = service.Run()
}
