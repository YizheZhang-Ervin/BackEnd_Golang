package util

import (
	"net/http"
	"time"
)

type HttpChecker struct {
	Servers    HttpServers
	FailMax    int
	RecovCount int //连续成功到达这个值，就会被标识为UP
}

func NewHtttpChecker(servers HttpServers) *HttpChecker {
	return &HttpChecker{Servers: servers, FailMax: 6, RecovCount: 3}
}
func (this *HttpChecker) Check(timeout time.Duration) {
	client := http.Client{}
	for _, server := range this.Servers {
		res, err := client.Head(server.Host)
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil { //宕机了
			this.Fail(server)
			continue
		}
		if res.StatusCode >= 200 && res.StatusCode < 400 {
			this.Success(server)
		} else {
			this.Fail(server)
		}
	}
}

func (this *HttpChecker) Fail(server *HttpServer) {
	if server.FailCount >= this.FailMax { //超过阈值
		server.Status = "DOWN"
	} else {
		server.FailCount++
	}
	server.SuccessCount = 0

}

func (this *HttpChecker) Success(server *HttpServer) {
	if server.FailCount > 0 {
		server.FailCount--
		server.SuccessCount++
		if server.SuccessCount == this.RecovCount {
			server.FailCount = 0
			server.Status = "UP"
			server.SuccessCount = 0
		}
	} else {
		server.Status = "UP"
	}

}
