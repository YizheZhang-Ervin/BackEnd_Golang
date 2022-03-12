// type 接口类型名 interface{
//     方法名1( 参数列表1 ) 返回值列表1
//     方法名2( 参数列表2 ) 返回值列表2
//     …
// }

package knowledges

import "fmt"

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

func interfaceDemo() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}
