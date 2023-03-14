package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/clientv3"
)

// 入口
func main() {
	// 连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect to etcd failed:", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// 设值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "xxKey", "xxVal")
	cancel()
	if err != nil {
		fmt.Println("put to etcd failed,err:", err)
		// 错误处理
		switch err {
			case context.Canceled:
				log.Fatalf("ctx is canceled by another routine: %v", err)
			case context.DeadlineExceeded:
				log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
			case rpctypes.ErrEmptyKey:
				log.Fatalf("client-side error: %v", err)
			default:
				log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
		return
	}

	// 取值
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "xxKey") //resp是相应对象
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed:", err)
		return
	}
	for _, ev := range resp.Kvs { //Kvs是响应对象的多个键值对
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

	// 监听
	//派一个哨兵 一直监视着key的变化（新增，修改，删除）
	ch := cli.Watch(context.Background(), "xxKey")
	//从通道中尝试取值（监视的信息）
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), evt.Kv.Value)
		}
	}
}
