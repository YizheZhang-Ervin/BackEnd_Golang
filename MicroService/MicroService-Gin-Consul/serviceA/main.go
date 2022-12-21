package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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
	r.GET("/api/add", func(c *gin.Context) {
		x, _ := strconv.Atoi(c.Query("x"))
		y, _ := strconv.Atoi(c.Query("y"))
		z := x + y
		c.String(200, fmt.Sprintf("z=%d", z))
	})
	return r
}

func main() {
	service := web.NewService(
		web.Name("serviceA"),
		web.Address(":50000"),
		web.Handler(InitWeb()),
		web.Registry(reg),
	)
	_ = service.Run()
}
