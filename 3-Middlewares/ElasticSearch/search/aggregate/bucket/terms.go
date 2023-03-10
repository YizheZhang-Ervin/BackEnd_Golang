package main
​
import (
    "context"
    "fmt"
    "github.com/olivere/elastic/v7"
    "log"
)
​
func main() {
    // 创建ES client
    client, err := elastic.NewClient()
    if err != nil {
        // Handle error
        panic(err)
    }
​
    // 执行ES请求需要提供一个上下文对象
    ctx := context.Background()
​
    // 创建Terms桶聚合
    aggs := elastic.NewTermsAggregation().
        Field("shop_id") // 根据shop_id字段值，对数据进行分组
​
    searchResult, err := client.Search().
        Index("shops"). // 设置索引名
        Query(elastic.NewMatchAllQuery()). // 设置查询条件
        Aggregation("shop", aggs). // 设置聚合条件，并为聚合条件设置一个名字
        Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
        Do(ctx) // 执行请求
​
    if err != nil {
        // Handle error
        panic(err)
    }
​
    // 使用Terms函数和前面定义的聚合条件名称，查询结果
    agg, found := searchResult.Aggregations.Terms("shop")
    if !found {
        log.Fatal("没有找到聚合数据")
    }
​
    // 遍历桶数据
    for _, bucket := range agg.Buckets {
        // 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
        bucketValue := bucket.Key
​
        // 打印结果， 默认桶聚合查询，都是统计文档总数
        fmt.Printf("bucket = %q 文档总数 = %d\n", bucketValue, bucket.DocCount)
    }
}