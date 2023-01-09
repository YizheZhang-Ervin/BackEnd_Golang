package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)

	//func Listen(network, address string) (Listener, error) {
	//net.Listen("tcp", ":8848") //简写，冒号前面默认是本机: 127.0.0.1
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	fmt.Println("监听中...")

	//Accept() (Conn, error)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}

	fmt.Println("连接建立成功!")

	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make来创建字节切片, byte ==> uint8

	//Read(b []byte) (n int, err error)
	//cnt：真正读取client发来的数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Client =====> Server, 长度:", cnt, "，数据:", string(buf[0:cnt]))

	//服务器对客户端请求进行响应 ,将数据转成大写 "hello" ==> HELLO
	//func ToUpper(s string) string {
	upperData := strings.ToUpper(string(buf[0:cnt]))

	//Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client  <====== Server, 长度:", cnt, "，数据:", upperData)

	//关闭连接
	_ = conn.Close()
}
