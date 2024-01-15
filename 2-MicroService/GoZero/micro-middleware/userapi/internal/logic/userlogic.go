package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/types/user"
	"time"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	userResponse, err := l.svcCtx.UserRpc.SaveUser(ctx, &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserLogic) GetUser(req *types.IdRequest) (interface{}, error) {
	value := l.ctx.Value("userId")
	logx.Infof("get token content: %s \n ", value)
	if req.Id == "1" {
		return nil, errors.New("id参数不正确")
	}
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserLogic) Login(t *types.LoginRequest) (string, error) {
	logx.Infof("login...")
	userId := 100
	auth := l.svcCtx.Config.Auth
	return l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
}

func (l *UserLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
