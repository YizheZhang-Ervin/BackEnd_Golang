package main

import (
	"fmt"
	"micro/AppLib"

	"gopkg.in/go-playground/validator.v9"
)

type Users struct {
	Username string `validate:"required"`
	Userpwd  string `validate:"required,min=6,max=18" vmsg:"验证密码必须6位以上"` //这里的vmsg是随便定义的，待会需要用我们自定义的函数去解析他，只要你高兴定义abc都可以
}

func main() {
	user := Users{Username: "xiahualou", Userpwd: "123"}
	valid := validator.New()   //新建一个验证器
	err := valid.Struct(&user) //传入结构体校验字段
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs { //errors是一个错误类型的切片
				fmt.Println(e.Value())
				fmt.Println(e.Field())
				fmt.Println(e.Tag())
				AppLib.GetValidMsg(user, e.Field()) //调用我们自定义的函数去输出更好的日志信息
			}
		}
	}
}
