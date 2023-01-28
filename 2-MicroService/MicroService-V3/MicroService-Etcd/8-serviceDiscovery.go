package util

import (
	"context"
	"regexp"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Client struct {
	client   *clientv3.Client
	Services []*ServiceInfo
}

type ServiceInfo struct {
	ServiceID   string
	ServiceName string
	ServiceAddr string
}

func NewClient() *Client {
	config := clientv3.Config{
		Endpoints:   []string{"106.12.72.181:23791", "106.12.72.181:23792"},
		DialTimeout: 10 * time.Second,
	}
	client, _ := clientv3.New(config)
	return &Client{client: client}
}

func (this *Client) GetService() {
	kv := clientv3.NewKV(this.client)
	res, _ := kv.Get(context.TODO(), "/services", clientv3.WithPrefix())
	for _, item := range res.Kvs {
		this.parseService(item.Key, item.Value)
	}
}

func (this *Client) LoadService() error {
	kv := clientv3.NewKV(this.client)
	res, err := kv.Get(context.TODO(), "/services", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, item := range res.Kvs {
		this.parseService(item.Key, item.Value)
	}
	return nil
}
func (this *Client) parseService(key []byte, value []byte) {
	reg := regexp.MustCompile("/services/(\\w+)/(\\w+)")
	if reg.Match(key) {
		idandname := reg.FindSubmatch(key) //idandname是一个二维切片
		sid := idandname[1]                //sid是一维切片，用string(sid)可以转成string
		sname := idandname[2]
		this.Services = append(this.Services, &ServiceInfo{ServiceID: string(sid),
			ServiceName: string(sname), ServiceAddr: string(value)})
	}
}
