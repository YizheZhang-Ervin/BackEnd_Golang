// type 类型名 struct {
//     字段名 字段类型
//     字段名 字段类型
//     …
// }
package knowledges

import "fmt"

//类型定义
type NewInt int

//类型别名
type MyInt = int

func typeDemo() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}

type person struct {
	name string
	city string
	age  int8
}

func structureDemo() {
	var p1 person
	p1.name = "xx"
	p1.city = "11"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={xx 11 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"xx", city:"11", age:18}
}

func anonymousStructureDemo() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}
