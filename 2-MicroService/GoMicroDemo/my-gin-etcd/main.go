package main

import (
	"fmt"
	"my-gin-etcd/handler"

	"my-gin-etcd/middlewares"
	"my-gin-etcd/utils"

	"github.com/gin-gonic/gin"
)

// Cors 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// 设置路由
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	// 不用认证的路由
	r.GET("/:key1/*key2", handler.GetHandler)
	// 需要认证的路由 -> post接口访问请求：
	r.Group("/auth", gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	{
		r.POST("/admin", handler.PostHandler)
	}
	return r
}

// 主函数启动
func main() {
	// 获取配置文件
	config := utils.ReadConfig("app", "json", "./config/")
	// 启动web服务器
	r := setupRouter()
	err := r.Run(config.GetString("ip") + ":" + config.GetString("port"))
	if err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	// 连接etcd
	middlewares.ConnectEtcd(config.GetString("etcd"))
}
