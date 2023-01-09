package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("cliet与server连接建立成功!")
	sendData := []byte("helloworld")

	// 如果用for循环保证多次向server发送数据
	// for {
	//向服务器发送数据
	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	fmt.Println("Client ===> Server cnt:", cnt, ", data:", string(sendData))

	//接收服务器返回的数据
	//创建buf，用于接收服务器返回的数据
	buf := make([]byte, 1024)
	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Client <==== Server , cnt:", cnt, ", data:", string(buf[0:cnt]))
	time.Sleep(1 * time.Second)
	// }

	_ = conn.Close()
}
