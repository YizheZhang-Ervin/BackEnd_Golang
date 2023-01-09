package routers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadApiA(e *gin.Engine) {
	e.Group("/apiA")
	{
		e.GET("/:key1/*key2", apiHandler)
		e.GET("/key3", urlHandler)
		e.GET("/req", reqHandler)
	}

}

// API参数
func apiHandler(context *gin.Context) {
	key1 := context.Param("key1") // XX/yy
	key2 := context.Param("key2") // xx/yy.zz
	context.JSON(http.StatusOK, gin.H{
		"message": key1 + "" + key2,
	})
}

// API参数2(URL数据解析)
func apiHandler2(c *gin.Context) {
	// 声明接收的变量
	var login LoginGet
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// URL参数
func urlHandler(context *gin.Context) {
	key3 := context.Query("key3")                       // XX?key=val
	key3Default := context.DefaultQuery("key3", "none") // XX?key=val
	context.JSON(http.StatusOK, gin.H{
		"message": key3 + "" + key3Default,
	})
}

// 获取请求头，请求体
func reqHandler(context *gin.Context) {
	body, _ := ioutil.ReadAll(context.Request.Body) // 请求体
	fmt.Println(body)
	for k, v := range context.Request.Header { // 请求头
		fmt.Println(k, v)
	}
	context.Writer.Header().Set("key", "value") // 响应头
	context.JSON(http.StatusOK, gin.H{
		"message": context.Request.RequestURI,
	})
}

// 定义接收数据的结构体
type LoginGet struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}
