package main

// 更新文档
_, err := client.Update().
        Index("blogs"). // 设置索引名
        Id("1"). // 文档id
        Doc(map[string]interface{}{"Title": "新的文章标题"}). // 更新Title="新的文章标题"，支持传入键值结构
        Do(ctx) // 执行ES查询
if err != nil {
   // Handle error
   panic(err)
}

// 根据条件更新文档
_, err = client.UpdateByQuery("blogs").
                // 设置查询条件，这里设置Author=tizi
        Query(elastic.NewTermQuery("Author", "tizi")).
                // 通过脚本更新内容，将Title字段改为1111111
        Script(elastic.NewScript( "ctx._source['Title']='1111111'")).
                // 如果文档版本冲突继续执行
        ProceedOnVersionConflict(). 
        Do(ctx)