package util

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Service struct {
	client *clientv3.Client
}

func NewService() *Service {
	config := clientv3.Config{
		Endpoints:   []string{"106.12.72.181:23791", "106.12.72.181:23792"},
		DialTimeout: 10 * time.Second,
	}
	client, _ := clientv3.New(config)
	return &Service{client: client}
}

//注册服务
func (this *Service) RegService(id string, name string, address string) error {
	kv := clientv3.NewKV(this.client)
	key_prefix := "/services/"
	ctx := context.Background()
	lease := clientv3.NewLease(this.client)
	leaseRes, err := clientv3.NewLease(this.client).Grant(ctx, 20) //设置租约过期时间为20秒
	if err != nil {
		return err
	}
	_, err = kv.Put(context.Background(), key_prefix+id+"/"+name, address, clientv3.WithLease(leaseRes.ID)) //把服务的key绑定到租约下面
	if err != nil {
		return err
	}
	keepaliveRes, err := lease.KeepAlive(context.TODO(), leaseRes.ID) //续租时间大概自动为租约的三分之一时间，context.TODO官方定义为是你不知道要传什么context的时候就用这个
	if err != nil {
		return err
	}
	go lisKeepAlive(keepaliveRes)
	return err
}

func lisKeepAlive(keepaliveRes <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case ret := <-keepaliveRes:
			if ret != nil {
				fmt.Println("续租成功", time.Now())
			}
		}
	}
}

//反注册服务
func (this *Service) UnregService(id string) error {
	kv := clientv3.NewKV(this.client)
	key_prefix := "/services/" + id
	_, err := kv.Delete(context.Background(), key_prefix, clientv3.WithPrefix())
	fmt.Println(err)
	return err
}
