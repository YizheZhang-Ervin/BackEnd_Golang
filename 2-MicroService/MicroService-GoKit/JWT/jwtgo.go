package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Uname              string `json:"username"`
	jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

func main() {
	sec := []byte("123abc")
	//hs256
	//生成jwt
	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{Uname: "xiahualou"}) //使用HS256加密算法加密
	token, _ := token_obj.SignedString(sec)                                               //把秘钥传进去，生成签名token
	fmt.Println(token)

	uc := UserClaim{}
	//验证jwt
	getToken, _ := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return sec, nil //这里是对称加密，所以只要有人拿到了这个sec就可以进行访问不安全
	})
	//用下面这种解析方式可以把解析后的结果保存到结构体中去
	getToken, _ = jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
		return sec, nil
	})
	//验证jwt是否有效
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Uname) //使用断言判断具体的claim直接取值
	}
}
