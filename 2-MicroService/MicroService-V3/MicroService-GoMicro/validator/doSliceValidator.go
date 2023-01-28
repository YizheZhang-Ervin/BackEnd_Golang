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
	//下面的tag中有一个特别的dive标签，作用是进到切片内部对元素进行校验，写在dive之前是对切片的校验，dive之后是对元素的校验
	Usertags []string `validate:"required,min=1,max=5,unique,dive,usertag" vmsg:"用户标签不合法"` //切片中min和max表示切片最大长度和最小长度，unique表示切片数据集不可以重复，usertag是绑定到自定义的正则验证用的
}

func main() {
	userTags := []string{"aa", "#b", "c", "d", "e"}
	user := &Users{Username: "shenyi", Userpwd: "123123", Testname: "wqeqdasd", Usertags: userTags}
	valid := validator.New()
	//加入自定义的正则验证tag
	err := AppLib.AddRegexTag("username", "[a-zA-Z]\\w{5,19}", valid)
	if err != nil {
		log.Fatal(err)
	}
	err = AppLib.AddRegexTag("usertag", "^[a-zA-Z0-9]{1,}", valid)
	if err != nil {
		log.Fatal(err)
	}
	err = AppLib.ValidErrMsg(user, valid.Struct(user)) //这里才是真正调用验证，包括我们设置的正则验证
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("验证成功")
}
