package Wrappers

import (
	"context"
	"go-micro/Services"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type ProdsWrapper struct { //官方提供的例子，创建自己的struct，嵌套go-micro的client
	client.Client
}

func defaultProds(rsp interface{}) { //将rsp中传入响应值，这里响应值是我们proto定义好的返回值
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(20+i, "prodname"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*Services.ProdListResponse) //类型断言为我们定义好的返回值
	result.Data = models
}

func newProd(i int32, s string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: i, ProdName: s}
}

//重写Call方法
func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint() //req.Service()是服务名.Endpoint是方法,这里是ProdService.GetProdsList,这个名字并不会对结果有影响，只是这里规范定义而已，其实定义hello world也可以运行
	/*
	   service ProdService{
	       rpc GetProdsList (ProdsRequest) returns (ProdListResponse);
	   }
	*/
	configA := hystrix.CommandConfig{
		Timeout: 5000,
	}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp) //调用rpc api接口
	}, func(e error) error { //降级函数
		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
