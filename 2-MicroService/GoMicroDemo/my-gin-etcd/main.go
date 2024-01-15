package main

import (
	"fmt"
	"net/http"

	"my-gin-etcd/myconfig"
	"my-gin-etcd/myetcd"

	"github.com/gin-gonic/gin"
)

// 要认证的post接口访问请求：
// Zm9vOmJhcg== is base64("foo:bar")
// curl -X POST http://localhost:8080/admin \
// -H 'authorization: Basic Zm9vOmJhcg==' \
// -H 'content-type: application/json' \
// -d '{"value":"bar"}'

// 设置路由
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	// 不用认证的路由
	r.GET("/:key1/*key2", getHandler)
	// 需要认证的路由
	r.Group("/auth", gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	{
		r.POST("/admin", postHandler)
	}
	return r
}

// get处理器
func getHandler(c *gin.Context) {
	key1 := c.Param("key1") // XX/yy
	key2 := c.Param("key2") // xx/yy.zz
	c.JSON(http.StatusOK, gin.H{
		"message": key1 + "" + key2,
	})
}

// post处理器
func postHandler(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	var json struct {
		Value string `json:"value" binding:"required"`
	}
	if c.Bind(&json) == nil {
		fmt.Println(user, json.Value)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// 主函数启动
func main() {
	// 获取配置文件
	config := myconfig.Connect("application", "json", "../configs/")
	// 启动web服务器
	r := setupRouter()
	err := r.Run(config.GetString("port"))
	if err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	// 连接etcd
	myetcd.Connect(config.GetString("etcd"))
}
