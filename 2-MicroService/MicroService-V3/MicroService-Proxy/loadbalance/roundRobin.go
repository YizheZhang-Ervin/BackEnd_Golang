package util

import "sort"

type HttpServer struct { //目标server类
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

type LoadBalance struct { //负载均衡类
	Servers  []*HttpServer
	CurIndex int //指向当前访问的服务器
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HttpServer, 0)}
}

func (this *LoadBalance) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}

func (this *LoadBalance) RoundRobin() *HttpServer {
	server := this.Servers[this.CurIndex]
	//this.CurIndex++
	//if this.CurIndex >= len(this.Servers) {
	//    this.CurIndex = 0
	//}
	this.CurIndex = (this.CurIndex + 1) % len(this.Servers) //因为一个数的余数永远在0-它本身之间，所以用这种方式得到的轮询更好
	return server
}

func (this *LoadBalance) RoundRobinByWeight() *HttpServer {
	server := this.Servers[ServerIndices[this.CurIndex]]
	this.CurIndex = (this.CurIndex + 1) % len(ServerIndices) //ServersIndices存放的是按照权重排放的索引，如3，1，2 则ServerIndices=[0,0,0,1,2,2] 然后遍历ServerIndices可以拿到按照权重得到的索引，在每一次遍历中用索引得到[3,1,2]对应的index，一个数的余数只能是0到他自己
	return server
}

func (this *LoadBalance) RoundRobinByWeight3() *HttpServer { //平滑加权轮询
	for _, s := range this.Servers {
		s.CWeight = s.CWeight + s.Weight
	}
	sort.Sort(this.Servers)
	max := this.Servers[0]

	max.CWeight = max.CWeight - SumWeight
	return max
}

var LB *LoadBalance
var ServerIndices []int

func init() {
	LB = NewLoadBalance()
	LB.AddServer(NewHttpServer("http://localhost:8001/web1", 5))  //web1
	LB.AddServer(NewHttpServer("http://localhost:8002/web2", 15)) //web2
	for index, server := range LB.Servers {
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndices = append(ServerIndices, index)
			}
		}
	}
}
