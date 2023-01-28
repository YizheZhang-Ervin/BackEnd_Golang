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
						return fmt.Errorf("%s", e)
					}
				} else {
					return fmt.Errorf("%s", e) //因为当切片中某个字段验证错误了，这时候e.Field()是Usertags[i],而针对Users这个结构体的字段没有Usertags[i]这个字段，所以上面的exist会是false，这时候如果我们不处理，返回值就是nil，验证就通过了，所以我们需要另外一个else额外处理一下
				}
			}
		}
	}
	return nil
}
