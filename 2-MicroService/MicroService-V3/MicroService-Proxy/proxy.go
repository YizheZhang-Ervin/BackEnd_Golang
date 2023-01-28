package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"regexp"
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
	for k, v := range util.ProxyConfigs {
		if match, _ := regexp.MatchString(k, r.URL.Path); match == true {
			target, _ := url.Parse(v)                           //v是目标网站地址
			proxy := httputil.NewSingleHostReverseProxy(target) //go内置的反向代理
			proxy.ServeHTTP(w, r)
			return
		}
	}
	w.Write([]byte("default index"))
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)

}
