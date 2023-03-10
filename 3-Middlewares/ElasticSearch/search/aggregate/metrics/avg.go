package main

// 创建Avg指标聚合
aggs := elastic.NewAvgAggregation().
        Field("price") // 设置统计字段
​
searchResult, err := client.Search().
        Index("kibana_sample_data_flights"). // 设置索引名
        Query(elastic.NewMatchAllQuery()). // 设置查询条件
        Aggregation("avg_price", aggs). // 设置聚合条件，并为聚合条件设置一个名字
        Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
        Do(ctx) // 执行请求
​
if err != nil {
    // Handle error
    panic(err)
}
​
// 使用Avg函数和前面定义的聚合条件名称，查询结果
agg, found := searchResult.Aggregations.Avg("avg_price")
if found {
    // 打印结果，注意：这里使用的是取值运算符
    fmt.Println(*agg.Value)
}