package main

// 创建创建一条微博
msg1 := Weibo{User: "olivere", Message: "打酱油的一天", Retweets: 0}
​
// 使用client创建一个新的文档
put1, err := client.Index().
        Index("weibo"). // 设置索引名称
        Id("1"). // 设置文档id
        BodyJson(msg1). // 指定前面声明的微博内容
        Do(ctx) // 执行请求，需要传入一个上下文对象
if err != nil {
        // Handle error
        panic(err)
    }
​
fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)