package main
​
import (
    "context"
    "fmt"
    "github.com/olivere/elastic/v7"
    "log"
    "os"
    "reflect"
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
        // 创建Client, 连接ES
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
    // 创建term查询条件，用于精确查询
    termQuery := elastic.NewTermQuery("Author", "tizi")
    
    searchResult, err := client.Search().
        Index("blogs").   // 设置索引名
        Query(termQuery).   // 设置查询条件
        Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
        From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
        Size(10).   // 设置分页参数 - 每页大小
        Pretty(true).       // 查询结果返回可读性较好的JSON格式
        Do(ctx)             // 执行请求
​
    if err != nil {
        // Handle error
        panic(err)
    }
​
    fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
​
​
    if searchResult.TotalHits() > 0 {
        // 查询结果不为空，则遍历结果
        var b1 Article
        // 通过Each方法，将es结果的json结构转换成struct对象
        for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
            // 转换成Article对象
            if t, ok := item.(Article); ok {
                fmt.Println(t.Title)
            }
        }
    }
}