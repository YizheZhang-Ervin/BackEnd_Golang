package logic

import (
	"context"

	"order/internal/svc"
	"order/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
