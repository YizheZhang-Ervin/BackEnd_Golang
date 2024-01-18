package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHandler get处理器
func GetHandler(c *gin.Context) {
	key1 := c.Param("key1") // XX/yy
	key2 := c.Param("key2") // xx/yy.zz
	c.JSON(http.StatusOK, gin.H{
		"message": key1 + "" + key2,
	})
}

// PostHandler post处理器
func PostHandler(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	var json struct {
		Value string `json:"value" binding:"required"`
	}
	if c.Bind(&json) == nil {
		fmt.Println(user, json.Value)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
