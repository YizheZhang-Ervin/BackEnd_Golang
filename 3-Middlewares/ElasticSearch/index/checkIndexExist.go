package main

// 检测下weibo索引是否存在
exists, err := client.IndexExists("weibo").Do(ctx)
if err != nil {
    // Handle error
    panic(err)
}