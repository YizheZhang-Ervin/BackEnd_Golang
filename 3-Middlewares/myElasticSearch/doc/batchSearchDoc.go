package main


// 查询id等于1,2,3的博客内容
result, err := client.MultiGet().
Add(elastic.NewMultiGetItem(). // 通过NewMultiGetItem配置查询条件
	Index("blogs"). // 设置索引名
	Id("1")). // 设置文档id
Add(elastic.NewMultiGetItem().Index("blogs").Id("2")).
Add(elastic.NewMultiGetItem().Index("blogs").Id("3")).
Do(ctx) // 执行请求

if err != nil {
panic(err)
}
​
// 遍历文档
for _, doc := range result.Docs {
// 转换成struct对象
var content Article
tmp, _ := doc.Source.MarshalJSON()
err := json.Unmarshal(tmp, &content)
if err != nil {
	panic(err)
}
​
fmt.Println(content.Title)
}