package main

// 根据id删除一条数据
_, err := client.Delete().
        Index("weibo").
        Id("1").
        Do(ctx)
if err != nil {
    // Handle error
    panic(err)
}