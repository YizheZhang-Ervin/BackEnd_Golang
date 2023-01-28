package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var r = rate.NewLimiter(1, 5) //1表示每次放进筒内的数量，桶内的令牌数是5，最大令牌数也是5，这个筒子是自动补充的，你只要取了令牌不管你取多少个，这里都会在每次取完后自动加1个进来，因为我们设置的是1

func MyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) { //这里使用http.HandlerFunc进行类型转换把匿名函数转换成了type http.HandlerFunc，因为HandlerFunc实现了ServeHttp方法所以是http.Handler的实例
		if !r.Allow() { //r.Allow()实际上是r.AllowN(time.Now(), 1)的缩写,这里每次拿一个令牌，拿一个存一个，如果有很多个线程来访问，一次把5个拿光了就会走这条,或者请求太快，也会一下子取光
			http.Error(writer, "too many requests", http.StatusTooManyRequests)
		} else {
			next.ServeHTTP(writer, request)
			/*处理完了之后，再调用next.ServerHTTP继续完成请求，因为HandlerFunc类型实现的ServeHTTP是这样调用的，也就是调用ServeHTTP方法其实最终还是调用到了我们这里的匿名函数，所以当没有超出频率限制时，仍然需要返回正确的结果给用户，直接调用next也就是mux的ServeHTTP方法，也就是最终调用了绑定路由的方法去返回结果
			  func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
			      f(w, r)
			  }

			*/
		}
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK!!!"))
	})
	http.ListenAndServe(":8080", MyLimit(mux))
}
