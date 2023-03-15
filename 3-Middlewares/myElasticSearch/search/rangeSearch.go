package main

// 例1 等价表达式： Created > "2020-07-20" and Created < "2020-07-29"
rangeQuery := elastic.NewRangeQuery("Created").
        Gt("2020-07-20").
        Lt("2020-07-29")
​
// 例2 等价表达式： id >= 1 and id < 10
rangeQuery := elastic.NewRangeQuery("id").
        Gte(1).
        Lte(10)