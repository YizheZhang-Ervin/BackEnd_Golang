package main

_, err := client.Update().
        Index("weibo"). // 设置索引名
        Id("1"). // 文档id
        Doc(map[string]interface{}{"retweets": 0}). // 更新retweets=0，支持传入键值结构
        Do(ctx) // 执行ES查询
if err != nil {
   // Handle error
   panic(err)
}