package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`                 //==> 在使用json编码时，这个编码不参与
	Subject string `json:"Subject_name"`      //==> 在json编码时，这个字段会编码程Subject_name
	Age     int    `json:"age,string"`        //==>在json编码时，将age转成程string类型, 一定要两个字段:名字,类型，中间不能有空格
	Address string `json:"address,omitempty"` //==》在json编码时，如果这个字段是空的，那么忽略掉，不参与编码

	//注意，gender是小写的, 小写字母开头的，在json编码时会忽略掉
	gender string
}

type Master struct {
	Name    string
	Subject string
	Age     int
	Address string
	gender  string
}

func main() {

	t1 := Teacher{
		Name:    "Duke",
		Subject: "Golang",
		Age:     18,
		gender:  "Man",
		Address: "北京",
	}

	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal(&t1)

	fmt.Println("encodeInfo:", string(encodeInfo))

	//解码
	t2 := Teacher{}
	_ = json.Unmarshal(encodeInfo, &t2)
	fmt.Println("t2:", t2.Subject)

	m1 := Master{}
	_ = json.Unmarshal(encodeInfo, &m1)
	fmt.Println("m1:", m1)

}
