package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求, 服务端在注册rpc对象是, 能让编译期检测出 注册对象是否合法.

// 创建接口, 在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

// 调用该方法时, 需要给 i 传参, 参数应该是 实现了 HelloWorld 方法的类对象!
func RegisterService(i MyInterface)  {
	rpc.RegisterName("hello", i)
}


// -----------------客户端用

// 向调用本地函数一样,调用远程函数.
// 定义类
type Myclient struct {
	c *rpc.Client
}

// 由于使用了 c 调用 Call, 因此需要初始化 c
func InitClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)

	return Myclient{c:conn}
}

// 实现函数, 原型参照上面的 Interface来实现.
func (this *Myclient) HelloWorld (a string, b *string) error {

	// 参数1, 参照上面的 Interface , RegisterName 而来.  a :传入参数  b:传出参数.
	return this.c.Call("hello.HelloWorld", a, b)
}
