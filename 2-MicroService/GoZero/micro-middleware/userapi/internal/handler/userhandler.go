package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"userapi/internal/logic"
	"userapi/internal/svc"
	"userapi/internal/types"
)

type UserHandler struct {
	serCtx *svc.ServiceContext
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req types.IdRequest
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewUserLogic(r.Context(), h.serCtx)
	resp, err := l.GetUser(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewUserLogic(r.Context(), h.serCtx)
	resp, err := l.Login(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func NewUserHandler(serverCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		serCtx: serverCtx,
	}
}
