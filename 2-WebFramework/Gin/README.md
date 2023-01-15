# Gin语法
```
（1）RESTful
（2）JSON序列化
（3）请求参数
（4）form参数
（5）JSON参数
（6）path参数
（7）参数绑定
（8）文件上传
（9）重定向
（10）路由
（11）中间件
```

# 1. RESTful接口
```
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
 
/**
  RESTful API开发实例
*/
func main() {
    engine := gin.Default()
    engine.GET("/user", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "message": "GET",
        })
    })
    engine.POST("/user", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "message": "POST",
        })
    })
    engine.PUT("/user", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "message": "PUT",
        })
    })
    engine.DELETE("/user", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "message": "DELETE",
        })
    })
    engine.Run(":8088")
}
```

# 2. 序列化
```
func main() {
    engine := gin.Default()
    engine.GET("/json1", func(context *gin.Context) {
        //手动拼接JSON串
        context.JSON(http.StatusOK, gin.H{"name": "张三", "age": 24})
    })
 
    engine.GET("/json2", func(context *gin.Context) {
        //使用结构体方式
        context.JSON(http.StatusOK, Student{"Pete", 12})
    })
 
    engine.Run("127.0.0.1:8088")
}
 
type Student struct {
    Name string `json:"name"`
    Age  int `json:"age"`
}
```

# 3. 请求参数
```
engine.GET("/search", func(context *gin.Context) {
    query := context.DefaultQuery("name", "默认值")
    age := context.Query("age")
    fmt.Printf("参数name:  %v  参数age： %v ", query, age)
    //接收到请求后再响应结果
    context.JSON(http.StatusOK, gin.H{"status": "OK", "msg": "message..."})
})
```

# 4. 表单
```
engine.POST("/search", func(context *gin.Context) {
    name := context.DefaultPostForm("name", "默认值")
    age := context.PostForm("age")
    fmt.Printf("参数name:  %v  参数age： %v ", name, age)
    //接收到请求后再响应结果
    context.JSON(http.StatusOK, gin.H{"status": "OK", "msg": "message..."})
})
```

# 5. JSON
```
engine.POST("/formJson", func(c *gin.Context) {
    //读取row格式请求体数据
    b, _ := c.GetRawData()
    // 定义map或结构体
    var m map[string]interface{}
    // 反序列化
    _ = json.Unmarshal(b, &m)
    fmt.Println("JSON请求参数", m)
    c.JSON(http.StatusOK, m)
})
```

# 6. 路径参数
```
//获取路径参数
engine.GET("/search/:name/:age", func(context *gin.Context) {
    name := context.Param("name")
    age:=context.Param("age")
    context.JSON(http.StatusOK, gin.H{"name": name, "age": age,"msg":"路径参数测试"})
})
```
# 7. 参数绑定
```
type UserInfo struct {
    UserName string `form:"userName" json:"userName" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}
 
func main() {
    engine := gin.Default()
    engine.GET("/login", func(context *gin.Context) {
        var userInfo UserInfo
        err := context.ShouldBind(&userInfo)
        if err == nil {
            fmt.Printf("login info:%#v\n", userInfo)
            context.JSON(http.StatusOK, gin.H{
                "userName": userInfo.UserName,
                "password": userInfo.Password,
            })
        } else {
            context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })
 
 
    engine.POST("/login", func(context *gin.Context) {
        var userInfo UserInfo
        err := context.ShouldBind(&userInfo)
        if err == nil {
            fmt.Printf("login info:%#v\n", userInfo)
            context.JSON(http.StatusOK, gin.H{
                "userName": userInfo.UserName,
                "password": userInfo.Password,
            })
        } else {
            context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })
     
    engine.Run(":8088")
}
```

# 8. 文件上传
```
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)
 
func main() {
    engine := gin.Default()
    //单个文件上传
    engine.POST("/upload", func(context *gin.Context) {
 
        file, err := context.FormFile("file")
        if err != nil {
            context.JSON(http.StatusInternalServerError, gin.H{
                "message": err.Error(),
            })
            return
        }
        dst := fmt.Sprintf("./src/file/%s", file.Filename)
        // 上传文件到指定的目录
        _ = context.SaveUploadedFile(file, dst)
        context.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("'%s' 文件已上传", file.Filename),
        })
    })
 
    //多文件上传
    engine.POST("/mulUpload", func(context *gin.Context) {
        // Multipart form
        form, _ := context.MultipartForm()
        files := form.File["file"]
        for _, file := range files {
            dst := fmt.Sprintf("./src/file/%s", file.Filename)
            // 上传文件到指定的目录
            _ = context.SaveUploadedFile(file, dst)
        }
        context.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("%d 多文件已上传", len(files)),
        })
    })
    engine.Run(":8088")
}
```

# 9. 重定向
```
// HTTP重定向
engine.GET("/test", func(context *gin.Context) {
    context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
})
 
 
// 路由重定向
engine.GET("/test", func(context *gin.Context) {
    // 指定重定向的URL
    context.Request.URL.Path = "/hello"
    engine.HandleContext(context)
})
engine.GET("/hello", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{"message": "hello"})
})
```

# 10. 路由
```
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
 
/**
  路由和路由组
*/
func main() {
    engine := gin.Default()
    //普通路由
    engine.GET("/demo", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{"msg": "demo"})
    })
    //匹配任意请求Method的路由
    engine.Any("/any", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{"msg": "any"})
    })
    //处理未匹配的路由
    engine.NoRoute(func(context *gin.Context) {
        context.JSON(http.StatusNotFound, gin.H{"msg": "NoRoute"})
    })
    //路由组
    userGroup:=engine.Group("/user")
    {
        userGroup.GET("/list", func(context *gin.Context) {
            context.JSON(http.StatusNotFound, gin.H{"msg": "user.list"})
        })
        userGroup.POST("/add", func(context *gin.Context) {
            context.JSON(http.StatusNotFound, gin.H{"msg": "user.add"})
        })
    }
    stuGroup:=engine.Group("/student")
    {
        stuGroup.GET("/list", func(context *gin.Context) {
            context.JSON(http.StatusNotFound, gin.H{"msg": "student.list"})
        })
        stuGroup.POST("/add", func(context *gin.Context) {
            context.JSON(http.StatusNotFound, gin.H{"msg": "student.add"})
        })
    }
    engine.Run(":8088")
}
```

11. 中间件
```
//统计请求耗时
func computeCost() gin.HandlerFunc {
    return func(context *gin.Context) {
        start := time.Now()
        context.Set("message", "Hello")
        context.Next()
        cost := time.Since(start)
        fmt.Printf("处理耗时%d：", cost)
    }
}
 
func main() {
    engine := gin.Default()
    //全局注册中间件
    engine.Use(computeCost())
    //下面的第二个参数就是一个中间件，可以传入多个
    engine.GET("/demo", func(context *gin.Context) {
        value, exists := context.Get("message")
        if exists {
            context.JSON(http.StatusOK, gin.H{"msg": value})
        }
    })
    engine.Run(":8088")
}
```