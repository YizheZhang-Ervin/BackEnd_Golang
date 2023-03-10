package main

// 根据id查询文档
get1, err := client.Get().
        Index("blogs"). // 指定索引名
        Id("1"). // 设置文档id
        Do(ctx) // 执行请求
if err != nil {
    // Handle error
    panic(err)
}
if get1.Found {
    fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
}
​
# 手动将文档内容转换成go struct对象
msg2 := Article{}
// 提取文档内容，原始类型是json数据
data, _ := get1.Source.MarshalJSON()
// 将json转成struct结果
json.Unmarshal(data, &msg2)
// 打印结果
fmt.Println(msg2.Title)