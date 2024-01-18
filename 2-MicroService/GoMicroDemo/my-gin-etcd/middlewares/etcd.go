package middlewares

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/clientv3"
)

var cliPtr *clientv3.Client

// ConnectEtcd 连接
func ConnectEtcd(connectStr string) {
	// 连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{connectStr},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect to etcd failed:", err)
		return
	}
	fmt.Println("connect to etcd success")
	// defer cli.Close()
	cliPtr = cli
}

// SetValue 设值
func SetValue(key string, value string) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// cancel()
	_, err := cliPtr.Put(ctx, key, value)
	if err != nil {
		// fmt.Println("put from etcd failed:", err)
		printError(err)
		return
	}
}

// GetValue 取值
func GetValue(key string) []map[string]string {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// cancel()
	resp, err := cliPtr.Get(ctx, key) //resp是相应对象

	if err != nil {
		fmt.Println("get from etcd failed:", err)
		return nil
	}
	var kvMapList []map[string]string
	for _, ev := range resp.Kvs { //Kvs是响应对象的多个键值对
		var kvMap map[string]string
		kvMap = make(map[string]string)
		kvMap[string(ev.Key)] = string(ev.Value)
		kvMapList = append(kvMapList, kvMap)
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return kvMapList
}

// Watch 监控
func Watch(key string) {
	// 监听
	//派一个哨兵 一直监视着key的变化（新增，修改，删除）
	ch := cliPtr.Watch(context.Background(), key)
	//从通道中尝试取值（监视的信息）
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), evt.Kv.Value)
		}
	}
}

// 打印错误
func printError(err error) {
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
}
