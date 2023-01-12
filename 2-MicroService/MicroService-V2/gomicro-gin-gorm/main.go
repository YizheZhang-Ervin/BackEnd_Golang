package main

import (
	"github.com/gin-gonic/gin"
	"bj38web/web/controller"
	"bj38web/web/model"
)

// 添加gin框架开发3步骤
func main()  {

	// 初始化 Redis 链接池
	model.InitRedis()

	// 初始化 MySQL 链接池
	model.InitDb()

	// 初始化路由
	router := gin.Default()

	router.Static("/home", "view")

	// 添加路由分组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
	}

	// 启动运行
	router.Run(":8080")
}
