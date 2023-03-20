package main

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func main() {
	router := gin.Default()

	// 初始化容器.
	store, _ := redis.NewStore(10, "tcp", "192.168.6.108:6379", "", []byte("my-microservice1"))

	// 设置临时session：会话结束时过期
	// store.Options(sessions.Options{MaxAge:0})

	// 使用session中间件
	// 参数：前端显示的cookie名称，session存储的容器
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(context *gin.Context) {
		// cookie
		// 设值：参数 name,value,maxAge,path,domain,secure(地址栏是否可以看),httpOnly
		// context.SetCookie("kk", "vv", 60*60, "", "", true, true)
		// context.SetCookie("kk", "vv", 0, "", "", true, true)
		// 取值
		// val,_:=context.Cookie("kk")
		// fmt.Println(val)

		// session
		s := sessions.Default(context)
		// 设置session (修改session时, 需要Save函数配合.否则不生效)
		//s.Set("kkk", "vvv")
		//s.Save()
		v := s.Get("kkk")
		fmt.Println("获取 Session:", v.(string))
		context.Writer.WriteString("测试 Session ...")
	})

	router.Run(":9999")
}
