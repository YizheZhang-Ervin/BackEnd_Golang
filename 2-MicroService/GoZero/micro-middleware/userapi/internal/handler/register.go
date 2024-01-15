package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"userapi/internal/logic"
	"userapi/internal/types"
)

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.serCtx)
	resp, err := l.Register(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
