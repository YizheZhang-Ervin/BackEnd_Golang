package util

import (
	"math/rand"
	"time"
)

type LoadBalance struct {
	Servers []*ServiceInfo
}

func NewloadBalance(servers []*ServiceInfo) *LoadBalance {
	return &LoadBalance{Servers: servers}
}
func (this *LoadBalance) getByRand(sname string) *ServiceInfo {
	tmp := make([]*ServiceInfo, 0)
	for _, service := range this.Servers {
		if service.ServiceName == sname {
			tmp = append(tmp, service)
		}
	}
	if len(tmp) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(tmp) - 1)
	return this.Servers[i]
}
