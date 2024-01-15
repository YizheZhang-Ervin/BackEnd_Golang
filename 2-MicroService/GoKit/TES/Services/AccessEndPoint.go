package Services

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
)

const secKey = "123abc" //秘钥
type UserClaim struct {
	Uname string `json:"username"`
	jwt.StandardClaims
}
type IAccessService interface {
	GetToken(uname string, upass string) (string, error)
}
type AccessService struct{}

func (this *AccessService) GetToken(uname string, upass string) (string, error) {
	if uname == "jerry" && upass == "123" {
		userinfo := &UserClaim{Uname: uname}
		userinfo.ExpiresAt = time.Now().Add(time.Second * 60).Unix() //设置60秒的过期时间
		token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, userinfo)
		token, err := token_obj.SignedString([]byte(secKey))
		return token, err
	}
	return "", fmt.Errorf("error uname and password")
}

type AccessRequest struct {
	Username string
	Userpass string
	Method   string
}
type AccessResponse struct {
	Status string
	Token  string
}

func AccessEndpoint(accessService IAccessService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(AccessRequest)
		result := AccessResponse{Status: "OK"}
		if r.Method == "POST" {
			token, err := accessService.GetToken(r.Username, r.Userpass)
			if err != nil {
				result.Status = "error:" + err.Error()
			} else {
				result.Token = token
			}
		}
		return result, nil
	}
}
