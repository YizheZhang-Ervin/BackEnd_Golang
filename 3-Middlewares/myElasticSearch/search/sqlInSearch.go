package main

termsQuery := elastic.NewTermsQuery("Author", "tizi", "tizi365")
​
searchResult, err := client.Search().
        Index("blogs").   // 设置索引名
        Query(termsQuery).   // 设置查询条件
        Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
        From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
        Size(10).   // 设置分页参数 - 每页大小
        Do(ctx)             // 执行请求