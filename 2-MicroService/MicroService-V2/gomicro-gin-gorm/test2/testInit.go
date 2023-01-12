package test2

import "fmt"

// 首字母小写函数, 包作用域, 不能跨包使用!
func init()  {
	fmt.Println("测试 init 函数 ....")
}
