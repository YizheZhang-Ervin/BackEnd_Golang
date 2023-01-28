package AppLib

import (
	"fmt"
	"reflect"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

//封装一个通用的正则方法，省去每次都要写下面这段很长的代码
func AddRegexTag(tagName string, pattern string, v *validator.Validate) error {
	return v.RegisterValidation(tagName, func(fl validator.FieldLevel) bool {
		m, _ := regexp.MatchString(pattern, fl.Field().String()) //返回bool和错误类型
		return m
	}, false) //为true则值为null也做判断，为false则不判断
}

func ValidErrMsg(obj interface{}, err error) error {
	getObj := reflect.TypeOf(obj) //获取tag要用TypeOf
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
					if value, ok := f.Tag.Lookup("vmsg"); ok { //查找tag中有没有我们自定义的错误消息，有就用自定义的没有就用默认的
						return fmt.Errorf("%s", value)
					} else {
						return fmt.Errorf("%s", e.Value())
					}
				}
			}
		}
	}
	return nil
}
