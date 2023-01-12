package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"fmt"
)

func main()  {
	router := gin.Default()

	// 初始化容器.
	store, _ := redis.NewStore(10, "tcp", "192.168.6.108:6379", "", []byte("bj38"))

	// 设置临时session
/*	store.Options(sessions.Options{
		MaxAge:0,
	})*/

	// 使用容器
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(context *gin.Context) {
		// 调用session, 设置session数据
		s := sessions.Default(context)
		// 设置session
/*		s.Set("itcast", "itheima")
		// 修改session时, 需要Save函数配合.否则不生效
		s.Save()*/

		v := s.Get("itcast")
		fmt.Println("获取 Session:", v.(string))



		context.Writer.WriteString("测试 Session ...")
	})

	router.Run(":9999")
}
