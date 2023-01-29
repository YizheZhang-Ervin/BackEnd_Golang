package util

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
)

type HttpServer struct { //目标server类
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

type LoadBalance struct { //负载均衡类
	Servers []*HttpServer
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HttpServer, 0)}
}

func (this *LoadBalance) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}

func (this *LoadBalance) SelectForRand() *HttpServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.Servers))     //这里因为权重表为15个1和5个0组成，所以产生0到19的随机数
	return this.Servers[ServerIndices[index]] //通过随机数的索引获得服务器索引进而获得地址
}

func (this *LoadBalance) SelectByIpHash(ip string) *HttpServer {
	index := int(crc32.ChecksumIEEE([]byte(ip))) % len(this.Servers) //通过取余永远index都不会大于this.servers的长度
	return this.Servers[index]
}

func (this *LoadBalance) SelectByWeight(ip string) *HttpServer { //加权随机算法
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ServerIndices))
	return this.Servers[index]
}

func (this *LoadBalance) SelectByWeightBetter(ip string) *HttpServer {
	rand.Seed(time.Now().UnixNano())
	sumList := make([]int, len(this.Servers)) //this.servers是服务器列表
	sum := 0
	for i := 0; i < len(this.Servers); i++ {
		sum += this.Servers[i].Weight //如果是5，7，9权重之和为5 12 21，分三个区间[0:5) [5:12) [12,21) 0-20的随机数落在哪个区间就代表当前随机是哪个权重
		sumList[i] = sum              //生成权重区间列表

	}
	_rand := rand.Intn(sum)
	for index, value := range sumList {
		if _rand < value { //因为sumList是递增的，而且长度等于this.Servers所以遍历它比较随机数落在哪个区间就可以得到当前的权重是哪个
			return this.Servers[index]
		}
	}
	return this.Servers[0]
}

var LB *LoadBalance
var ServerIndices []int

func init() {
	LB := NewLoadBalance()
	LB.AddServer(NewHttpServer("http://localhost:8001/web1", 5))  //web1
	LB.AddServer(NewHttpServer("http://localhost:8002/web2", 15)) //web2
	for index, server := range LB.Servers {
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndices = append(ServerIndices, index)
			}
		}
	}
	fmt.Println(ServerIndices)
}
