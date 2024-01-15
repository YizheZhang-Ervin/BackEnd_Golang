package middlewares

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (*UserMiddleware) RegAndLoginHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("login 和 register 前面执行")
		next(w, r)
		logx.Info("login 和 register 后面面执行")
	}
}

func (*UserMiddleware) Global(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("all前面执行")
		next(w, r)
		logx.Info("all后面面执行")
	}
}
