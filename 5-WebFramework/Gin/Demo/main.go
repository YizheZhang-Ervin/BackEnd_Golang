package main

import (
	"Demo/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	port := ":9999"
	// 初始化创建
	r := gin.Default()
	// 处理跨域
	r.Use(Cors())
	// 导入路由
	routers.LoadApiA(r)
	routers.LoadApiB(r)
	// 启动
	if err := r.Run(port); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
