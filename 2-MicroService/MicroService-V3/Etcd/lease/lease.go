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
	fmt.Println(22)
	kv := clientv3.NewKV(this.client)
	key_prefix := "/services/"
	ctx := context.Background()
	leaseRes, err := clientv3.NewLease(this.client).Grant(ctx, 20) //设置租约过期时间为20秒
	if err != nil {
		return err
	}
	_, err = kv.Put(context.Background(), key_prefix+id+"/"+name, address, clientv3.WithLease(leaseRes.ID)) //把服务的key绑定到租约下面
	fmt.Println(err)
	return err
}

//反注册服务
func (this *Service) UnregService(id string) error {
	kv := clientv3.NewKV(this.client)
	key_prefix := "/services/" + id
	_, err := kv.Delete(context.Background(), key_prefix, clientv3.WithPrefix())
	fmt.Println(err)
	return err
}
