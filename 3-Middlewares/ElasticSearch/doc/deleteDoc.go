package main

// 根据id删除一条数据
_, err := client.Delete().
        Index("blogs").
        Id("1").  // 文档id
        Do(ctx)
if err != nil {
    // Handle error
    panic(err)
}

// 根据条件删除文档
_, _ = client.DeleteByQuery("blogs"). // 设置索引名
        // 设置查询条件为: Author = tizi
        Query(elastic.NewTermQuery("Author", "tizi")).
        // 文档冲突也继续删除
        ProceedOnVersionConflict().
        Do(ctx)