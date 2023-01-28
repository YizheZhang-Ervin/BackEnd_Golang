package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	. "反向代理/util"
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
	if r.URL.Path == "/a" {
		log.Println(r.RemoteAddr)

		newreq, _ := http.NewRequest(r.Method, "http://localhost:9091", r.Body)
		CLoneHead(r.Header, newreq.Header)                           //将请求头写入到模拟的请求头中去
		newreq.Header.Add("x-forwarded-for", r.RemoteAddr)           //将用户ip写入到代理请求头中去
		newreq.Header.Add("Authorization", "Basic c2hlbnlpOjEyMw==") //把BasicAuth的用户名和密码添加到header中，这样才可以过被代理的ip的验证

		newresponse, _ := http.DefaultClient.Do(newreq)
		getHeader := w.Header()
		CLoneHead(newresponse.Header, getHeader) //拷贝响应头 给客户端
		w.WriteHeader(newresponse.StatusCode)    // 写入http status

		defer newresponse.Body.Close()
		res_cont, _ := ioutil.ReadAll(newresponse.Body)
		w.Write(res_cont) // 写入响应给客户端
		return
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
