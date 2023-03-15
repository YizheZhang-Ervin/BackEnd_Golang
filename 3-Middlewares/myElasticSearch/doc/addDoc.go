package main
​
import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/olivere/elastic/v7"
    "log"
    "os"
    "time"
)
​
type Article struct {
    Title   string    // 文章标题
    Content string    // 文章内容
    Author  string    // 作者
    Created time.Time // 发布时间
}
​
​
func main() {
        // 创建client连接ES
    client, err := elastic.NewClient(
        // elasticsearch 服务地址，多个服务地址使用逗号分隔
        elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
        // 基于http base auth验证机制的账号和密码
        elastic.SetBasicAuth("user", "secret"),
        // 启用gzip压缩
        elastic.SetGzip(true),
        // 设置监控检查时间间隔
        elastic.SetHealthcheckInterval(10*time.Second),
        // 设置请求失败最大重试次数
        elastic.SetMaxRetries(5),
        // 设置错误日志输出
        elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
        // 设置info日志输出
        elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
​
    if err != nil {
        // Handle error
        fmt.Printf("连接失败: %v\n", err)
    } else {
        fmt.Println("连接成功")
    }
​
    // 执行ES请求需要提供一个上下文对象
    ctx := context.Background()
​
    // 定义一篇博客
    blog := Article{Title:"golang es教程", Content:"go如何操作ES", Author:"tizi", Created:time.Now()}
​
    // 使用client创建一个新的文档
    put1, err := client.Index().
        Index("blogs"). // 设置索引名称
        Id("1"). // 设置文档id
        BodyJson(blog). // 指定前面声明struct对象
        Do(ctx) // 执行请求，需要传入一个上下文对象
    if err != nil {
        // Handle error
        panic(err)
    }
​
    fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
}