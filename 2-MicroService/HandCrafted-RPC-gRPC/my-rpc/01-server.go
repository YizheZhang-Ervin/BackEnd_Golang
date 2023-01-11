package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法
func (this *World) HelloWorld (name string, resp *string) error {
	*resp = name + " 你好!"
	return nil
	//return errors.New("未知的错误!")
}

func main()  {

	// 1. 注册RPC服务, 绑定对象方法
	RegisterService(new(World))
/*	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册 rpc 服务失败!", err)
		return
	}*/

	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听 ...")
	// 3. 建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("链接成功...")

	// 4. 绑定服务
	jsonrpc.ServeConn(conn)
	//rpc.ServeConn(conn)
}
