package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	consulapi "github.com/hashicorp/consul/api"
)

var ConsulClient *consulapi.Client
var ServiceID string
var ServiceName string
var ServicePort int

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.3.14:8500"
	client, err := consulapi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	ServiceID = "userservice" + uuid.New().String()
	//因为最终这段代码是在不同的机器上跑的，是分布式的，有好几台机器提供相同的server，所以这里存到consul中的id必须是唯一的，否则只有一台服务器可以注册进去，这里使用uuid保证唯一性
}

func SetServiceNameAndPort(name string, port int) {
	ServiceName = name
	ServicePort = port
}

func RegService() {
	reg := consulapi.AgentServiceRegistration{}
	// reg.Name = "userservice"     //注册service的名字
	// reg.Address = "192.168.3.14" //注册service的ip
	// reg.Port = 8080              //注册service的端口
	reg.ID = ServiceID        //设置不同的Id，即使是相同的service name也得有不同的id
	reg.Name = ServiceName    //注册service的名字
	reg.Address = "localhost" //注册service的ip
	reg.Port = ServicePort    //注册service的端口
	fmt.Println(ServicePort)
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{} //创建consul的检查器
	check.Interval = "5s"                  //设置consul心跳检查时间间隔
	// check.HTTP = "http://192.168.3.14:8080/health" //设置检查使用的url
	check.HTTP = "http://192.168.3.14:" + strconv.Itoa(ServicePort) + "/health" //设置检查使用的url
	fmt.Println(check.HTTP)
	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}

}

func UnRegService() {
	ConsulClient.Agent().ServiceDeregister("userservice")
}
