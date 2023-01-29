package main

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"106.12.72.181:23791", "106.12.72.181:23792"},
		DialTimeout: 10 * time.Second,
	}
	client, _ := clientv3.New(config)
	defer client.Close()
	kv := clientv3.NewKV(client)
	ctx := context.Background()            //需要放入一个context，看自己需求选择合适的ctx
	kv.Put(ctx, "/services/user", "user1") //插入一条数据
}
