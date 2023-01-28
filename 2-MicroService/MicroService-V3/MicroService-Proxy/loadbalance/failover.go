package util

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"sort"
	"time"
)

type HttpServer struct { //目标server类
	Host         string
	Weight       int
	CWeight      int    //当前权重
	Status       string //健康检查
	FailCount    int    //计数器，默认是0
	SuccessCount int    //检查到连续成功，当连续成功的次数达到这个值，把宕机的的机器的FailCount立刻重置为0，加快服务器启动速度
}

type HttpServers []*HttpServer

func (p HttpServers) Len() int           { return len(p) }
func (p HttpServers) Less(i, j int) bool { return p[i].CWeight > p[j].CWeight }
func (p HttpServers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight, CWeight: 0}
}

type LoadBalance struct { //负载均衡类
	Servers  HttpServers
	CurIndex int //指向当前访问的服务器
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HttpServer, 0)}
}

func (this *LoadBalance) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}

func (this *LoadBalance) SelectForRand() *HttpServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.Servers))
	fmt.Println(index)
	return this.Servers[index]
}

func (this *LoadBalance) SelectByIpHash(ip string) *HttpServer {
	index := int(crc32.ChecksumIEEE([]byte(ip))) % len(this.Servers) //通过取余永远index都不会大于this.servers的长度
	return this.Servers[index]
}

func (this *LoadBalance) SelectByWeight(ip string) *HttpServer { //加权随机算法
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ServerIndices)) //这里因为权重表为15个1和5个0组成，所以产生0到19的随机数
	fmt.Println(this.Servers[ServerIndices[index]])
	return this.Servers[ServerIndices[index]] //通过随机数的索引获得服务器索引进而获得地址
}

func (this *LoadBalance) SelectByWeightBetter(ip string) *HttpServer {
	rand.Seed(time.Now().UnixNano())
	sumList := make([]int, len(this.Servers))
	sum := 0
	for i := 0; i < len(this.Servers); i++ {
		sum += this.Servers[i].Weight
		sumList[i] = sum

	}
	_rand := rand.Intn(sum)
	for index, value := range sumList {
		if _rand < value {
			return this.Servers[index]
		}
	}
	return this.Servers[0]
}

func (this *LoadBalance) RoundRobin() *HttpServer {
	server := this.Servers[this.CurIndex]
	//this.CurIndex ++
	//if this.CurIndex >= len(this.Servers) {
	//    this.CurIndex = 0
	//}
	this.CurIndex = (this.CurIndex + 1) % len(this.Servers)
	if server.Status == "Down" { //如果当前节点宕机了，则递归查找可以用的服务器
		return this.RoundRobin()
	}
	return server
}

func (this *LoadBalance) RoundRobinByWeight() *HttpServer {
	server := this.Servers[ServerIndices[this.CurIndex]]
	this.CurIndex = (this.CurIndex + 1) % len(ServerIndices)
	return server
}

func (this *LoadBalance) RoundRobinByWeight2() *HttpServer { //加权轮询 ,使用区间算法
	server := this.Servers[0]
	sum := 0
	//3:1:1
	for i := 0; i < len(this.Servers); i++ {
		sum += this.Servers[i].Weight //第一次是3   [0,3)  [3,4)   [4,5)
		if this.CurIndex < sum {
			server = this.Servers[i]
			if this.CurIndex == sum-1 && i != len(this.Servers)-1 {
				this.CurIndex++
			} else {
				this.CurIndex = (this.CurIndex + 1) % sum //这里是重要的一步
			}
			fmt.Println(this.CurIndex)
			break
		}
	}
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
var SumWeight int

func checkServers(servers HttpServers) {
	t := time.NewTicker(time.Second * 3)
	check := NewHtttpChecker(servers)
	for {
		select {
		case <-t.C:
			check.Check(time.Second * 2)
			for _, s := range servers {
				fmt.Println(s.Host, s.Status, s.FailCount)
			}
			fmt.Println("---------------------------------")
		}
	}
}

func init() {
	LB = NewLoadBalance()
	LB.AddServer(NewHttpServer("http://localhost:12346", 3)) //web1
	LB.AddServer(NewHttpServer("http://localhost:12347", 1)) //web2
	LB.AddServer(NewHttpServer("http://localhost:12348", 1)) //web2
	for index, server := range LB.Servers {
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndices = append(ServerIndices, index)
			}
		}
		SumWeight = SumWeight + server.Weight //计算加权总和
	}
	go checkServers(LB.Servers)

	//fmt.Println(ServerIndices)
}
