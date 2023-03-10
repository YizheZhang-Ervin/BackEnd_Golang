package main

// 创建ES client
client, err := elastic.NewClient()
if err != nil {
    // Handle error
    panic(err)
}
​
// 创建一个terms聚合，根据user字段分组，同时设置桶排序条件为按计数倒序排序，并且返回前10条桶数据
timeline := elastic.NewTermsAggregation().Field("user").Size(10).OrderByCountDesc()
// 创建Date histogram聚合,根据created时间字段分组，按年分组
histogram := elastic.NewDateHistogramAggregation().Field("created").CalendarInterval("year")
​
// 设置timeline的嵌套聚合条件，整体意思就是：首先按user字段分组，然后分组数据内，再次根据created时间字段按年分组，进行了两次分组。
timeline = timeline.SubAggregation("history", histogram)
​
// 执行ES查询
searchResult, err := client.Search().
    Index("twitter").                  // 设置索引名
    Query(elastic.NewMatchAllQuery()). // 设置查询条件
    Aggregation("timeline", timeline). // 设置聚合条件，并为聚合条件设置一个名字
    Pretty(true).                      // 返回可读的json格式
    Do(context.Background())           // 执行
if err != nil {
    // Handle error
    panic(err)
}
​
// 遍历ES查询结果，因为我们首先使用的是terms聚合条件，
// 所以查询结果先使用Terms函数和聚合条件的名字读取结果。
agg, found := searchResult.Aggregations.Terms("timeline")
if !found {
    // 没有查询到terms聚合结果
    log.Fatalf("we should have a terms aggregation called %q", "timeline")
}
​
// 遍历桶数据
for _, userBucket := range agg.Buckets {
    // 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
    user := userBucket.Key
​
    // 查询嵌套聚合查询的数据
    // 因为我们使用的是Date histogram聚合，所以需要使用DateHistogram函数和聚合名字获取结果
    histogram, found := userBucket.DateHistogram("history")
    if found {
        // 如果找到Date histogram聚合结果，则遍历桶数据
        for _, year := range histogram.Buckets {
            var key string
            if s := year.KeyAsString; s != nil {
                // 因为返回的是指针类型，这里做一下取值运算
                key = *s
            }
            // 打印结果
            fmt.Printf("user %q has %d tweets in %q\n", user, year.DocCount, key)
        }
    }
}