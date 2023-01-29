package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	"os"
	"os/signal"
	"反向代理/util"
)

type ProxyHandler struct {
}

func (*ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
		}
	}()

	url, _ := url2.Parse(util.LB.SelectByWeight(r.RemoteAddr).Host) //调用加权随机算法
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)

}
