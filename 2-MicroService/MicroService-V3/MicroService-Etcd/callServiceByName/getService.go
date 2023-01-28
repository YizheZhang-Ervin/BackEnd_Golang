package util

import (
	"context"
	"io/ioutil"
	"net/http"
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
		idandname := reg.FindSubmatch(key)
		sid := idandname[1]
		sname := idandname[2]
		this.Services = append(this.Services, &ServiceInfo{ServiceID: string(sid),
			ServiceName: string(sname), ServiceAddr: string(value)})
	}
}

//定义一个高阶函数返回处理具体url的函数
func (this *Client) GetService(sname string, method string, encodeFunc EncodeRequestFunc) Endpoint {
	for _, service := range this.Services {
		if service.ServiceName == sname {
			return func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error) {
				httpClient := &http.Client{}
				httpRequest, err := http.NewRequest(method, "http://"+service.ServiceAddr, nil)
				if err != nil {
					return nil, err
				}
				err = encodeFunc(ctx, httpRequest, requestParam) //这一步是关键，将httpRequest的url修改了
				if err != nil {
					return nil, err
				}
				res, err := httpClient.Do(httpRequest)
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return err, nil
				}
				return string(body), nil
			}
		}
	}
	return func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error) {
		return nil, err
	}
}
