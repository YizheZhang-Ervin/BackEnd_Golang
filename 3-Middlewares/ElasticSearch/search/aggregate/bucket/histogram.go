package main

/ 创建Histogram桶聚合
aggs := elastic.NewHistogramAggregation().
        Field("price"). // 根据price字段值，对数据进行分组
        Interval(50) //  分桶的间隔为50，意思就是price字段值按50间隔分组
​
searchResult, err := client.Search().
        Index("order"). // 设置索引名
        Query(elastic.NewMatchAllQuery()). // 设置查询条件
        Aggregation("prices", aggs). // 设置聚合条件，并为聚合条件设置一个名字
        Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
        Do(ctx) // 执行请求
​
if err != nil {
    // Handle error
    panic(err)
}
​
// 使用Histogram函数和前面定义的聚合条件名称，查询结果
agg, found := searchResult.Aggregations.Histogram("prices")
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