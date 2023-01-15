package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

// 创建中间件
func Test1(ctx *gin.Context)  {
	fmt.Println("1111")

	t := time.Now()

	ctx.Next()

	fmt.Println(time.Now().Sub(t))
}

// 创建 另外一种格式的中间件.
func Test2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("3333")

		context.Abort()

		fmt.Println("5555")
	}
}
func main()  {
	router := gin.Default()

	// 使用中间件
	router.Use(Test1)
	router.Use(Test2())

	router.GET("/test", func(context *gin.Context) {
		fmt.Println("2222")
		context.Writer.WriteString("hello world!")
	})

	router.Run(":9999")
}
