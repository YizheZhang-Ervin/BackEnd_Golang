package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct { //这个结构体主要是用来宣示当前公钥的使用者是谁，只有使用者和公钥的签名者是同一个人才可以用来正确的解密，还可以设置其他的属性，可以去百度一下
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

	token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, UserClaim{Uname: "xiahualou"}) //所有人给xiahualou发送公钥加密的数据，但是只有xiahualou本人可以使用私钥解密
	token, _ := token_obj.SignedString(priKey)

	uc := &UserClaim{}
	getToken, _ := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (i interface{}, e error) { //使用私钥解密
		return pubKey, nil //这里的返回值必须是公钥，不然解密肯定是失败
	})
	if getToken.Valid { //服务端验证token是否有效
		fmt.Println(getToken.Claims.(*UserClaim).Uname)
	}

}
