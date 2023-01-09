package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadApiB(e *gin.Engine) {
	e.Group("/apiB")
	{
		e.POST("/form", formHandler)
		e.POST("/form2", formHandler2)
		e.POST("/singlefile", singleFileHandler)
		e.POST("/multifile", multiFileHandler)
		e.POST("/json", jsonHandler)
	}
	// 设置上传文件的大小
	e.MaxMultipartMemory = 8 << 20
}

// 表单处理
func formHandler(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

// 表单处理2(表单数据解析)
func formHandler2(c *gin.Context) { // 声明接收的变量
	var form LoginPost
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 传单个文件
func singleFileHandler(c *gin.Context) {
	// input的name为xxfile
	file, err := c.FormFile("xxfile")
	if err != nil {
		c.String(500, "Upload Error")
	}
	c.SaveUploadedFile(file, file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": file.Filename})
}

// 传多个文件
func multiFileHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}

// json数据解析
func jsonHandler(c *gin.Context) {
	// 声明接收的变量
	var json LoginPost
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 定义接收数据的结构体
type LoginPost struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}
