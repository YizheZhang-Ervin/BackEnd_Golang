package main

// 1 must
// 创建bool查询
boolQuery := elastic.NewBoolQuery().Must()
​
// 创建term查询
termQuery := elastic.NewTermQuery("Author", "tizi")
matchQuery := elastic.NewMatchQuery("Title", "golang es教程")
​
// 设置bool查询的must条件, 组合了两个子查询
// 表示搜索匹配Author=tizi且Title匹配"golang es教程"的文档
boolQuery.Must(termQuery, matchQuery)
​
searchResult, err := client.Search().
        Index("blogs").   // 设置索引名
        Query(boolQuery).   // 设置查询条件
        Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
        From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
        Size(10).   // 设置分页参数 - 每页大小
        Do(ctx)             // 执行请求

// 2 must not 
// 创建bool查询
boolQuery := elastic.NewBoolQuery().Must()
​
// 创建term查询
termQuery := elastic.NewTermQuery("Author", "tizi")
​
// 设置bool查询的must not条件
boolQuery.MustNot(termQuery)

// 3 should
// 创建bool查询
boolQuery := elastic.NewBoolQuery().Must()
​
// 创建term查询
termQuery := elastic.NewTermQuery("Author", "tizi")
matchQuery := elastic.NewMatchQuery("Title", "golang es教程")
​
// 设置bool查询的should条件, 组合了两个子查询
// 表示搜索Author=tizi或者Title匹配"golang es教程"的文档
boolQuery.Should(termQuery, matchQuery)
