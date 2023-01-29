package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Uname              string `json:"username"`
	jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

func main() {
	priBytes, err := ioutil.ReadFile("./pem/private.pem")
	if err != nil {
		log.Fatal("私钥文件读取失败")
	}

	pubBytes, err := ioutil.ReadFile("./pem/public.pem")
	if err != nil {
		log.Fatal("公钥文件读取失败")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatal("公钥文件不正确")
	}

	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(priBytes)
	if err != nil {
		log.Fatal("私钥文件不正确")
	}
	user := UserClaim{Uname: "xiahualou"}
	user.ExpiresAt = time.Now().Add(time.Second * 5).Unix()      //UserClaim嵌套了jwt.StandardClaims，使用它的Add方法添加过期时间是5秒后，这里要使用unix()
	token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, user) //所有人给xiahualou发送公钥加密的数据，但是只有xiahualou本人可以使用私钥解密
	token, _ := token_obj.SignedString(priKey)
	//通过一秒一次for循环来验证过期生效
	for {
		uc := UserClaim{}
		getToken, err := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) { //使用私钥解密
			return pubKey, nil //这里的返回值必须是公钥，不然解密肯定是失败
		})
		if getToken.Valid { //服务端验证token是否有效
			fmt.Println(getToken.Claims.(*UserClaim).Uname)
		} else if ve, ok := err.(*jwt.ValidationError); ok { //官方写法招抄就行
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期或未启用")
			} else {
				fmt.Println("无法处理这个token", err)
			}

		}
		time.Sleep(time.Second)
	}

}
