package logic

import (
	"context"
	"strconv"

	"rpc-common/types/user"
	"user/internal/model"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	id, _ := strconv.ParseInt(in.Id, 10, 64)
	userData, err := l.svcCtx.UserRepo.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   userData.Name,
		Gender: userData.Gender,
	}, nil
}

func (l *GetUserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	data := &model.User{
		Name:   in.GetName(),
		Gender: in.GetGender(),
	}
	err := l.svcCtx.UserRepo.Save(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     strconv.FormatInt(data.Id, 10),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}
