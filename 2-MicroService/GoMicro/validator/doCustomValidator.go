package main

import (
	"fmt"
	"log"
	"micro/AppLib"

	"gopkg.in/go-playground/validator.v9"
)

type Users struct {
	Username string `validate:"required,min=6,max=20" vmsg:"用户名必须6位以上"`
	Userpwd  string `validate:"required,min=6,max=18" vmsg:"用户密码必须6位以上"`
	Testname string `validate:"username"  vmsg:"用户名规则不正确"` //这里的username对应v.RegisterValidation(tagName中的tagName，随便写写abc都可以但是要和它对应起来
}

func main() {

	user := &Users{Username: "shenyi", Userpwd: "123123", Testname: "wqeqdasd"}
	valid := validator.New()
	//加入自定义的正则验证tag
	err := AppLib.AddRegexTag("username", "[a-zA-Z]\\w{5,19}", valid)
	if err != nil {
		log.Fatal(err)
	}
	err = AppLib.ValidErrMsg(user, valid.Struct(user))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("验证成功")
}
