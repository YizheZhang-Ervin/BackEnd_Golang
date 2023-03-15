package main

// 创terms桶聚合
aggs := elastic.NewTermsAggregation().Field("shop_id")
// 创建Sum指标聚合
sumAggs := elastic.NewSumAggregation().Field("price")
// terms聚合嵌套指标聚合
aggs.SubAggregation("total_price", sumAggs)