package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//将所有的代码写在一个文件中，不做代码整理

type User struct {
	//名字
	name string
	//唯一的id
	id string
	//管道
	msg chan string
}

//创建一个全局的map结构，用户保存所有的用户
var allUsers = make(map[string]User)

//定义一个message全局通道，用于接收任何人发送过来消息
var message = make(chan string, 10)

func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//启动全局唯一的go程，负责监听message通道，写给所有的用户
	go broadcast()

	fmt.Println("服务器启动成功!")

	for {
		fmt.Println("=====> 主go程监听中...")

		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		//建立连接
		fmt.Println("建立连接成功!")

		//启动处理业务的go程
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn) {
	fmt.Println("启动业务...")

	//客户端与服务器建立连接的时候，会有ip和port ==> 当成user的id
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("clientAddr:", clientAddr)
	//创建user
	newUser := User{
		id:   clientAddr,            //id，我们不会修改，这个作为在map中的key
		name: clientAddr,            //可以修改，会提供rename命令修改，建立连接时，初始值与id相同
		msg:  make(chan string, 10), //注意需要make空间，否则无法写入数据
	}
	//添加user到map结构
	allUsers[newUser.id] = newUser

	//定义一个退出信号，用于监听client退出
	var isQuit = make(chan bool)

	//创建一个用于重置计数器的管道，用于告知watch函数，当前用户正在输入
	var restTimer = make(chan bool)

	//启动go程，负责监听退出信号
	go watch(&newUser, conn, isQuit, restTimer)

	//启动go程，负责将msg信息返回给客户端
	go writeBackToClient(&newUser, conn)

	//向message写入数据， 当前用户上线的消息，用于通知所有人（广播）
	loginInfo := fmt.Sprintf("[%s]:[%s] ===> 上线了login!!\n", newUser.id, newUser.name)
	message <- loginInfo

	for {
		//具体业务逻辑
		buf := make([]byte, 1024)

		//读取客户端发送过来的请求数据
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("客户端主动关闭ctrl + c，准备退出, err:", err)
			//map删除，用户， conn close掉
			//服务器还可以主动的退出
			//在这里不进行真正的退出动作，而是发送一个退出信号，统一做退出处理，可以使用新的管道来做信号传递
			isQuit <- true
		}

		if err != nil {
			fmt.Println("conn.Read err:", err, ", cnt:", cnt)
			return
		}

		fmt.Println("服务器接收客户端发送过来的数据为: ", string(buf[:cnt-1]), ", cnt:", cnt)
		//-------- 业务逻辑处理  开始----------
		//1. 查询当前所有的用户 who
		//	a. 先判断接收的数据是不是who  ==> 长度&&字符串
		userInput := string(buf[:cnt-1]) //这是用户输入的数据,最后一个是回车，我们去掉它
		// netcat 127.0.0.1 8080 输入 \who
		if len(userInput) == 4 && userInput == "\\who" {
			//  b. 遍历allUsers这个map (key : userid  value: user本身)，将id和name拼接成一个字符串，返回给客户端
			fmt.Println("用户即将查询所有用户信息!")

			//这个切片包含所有的用户信息
			var userInfos []string

			//[]string{"userid:z3, username:z3", "userid:l4, username:l4","userid:w5, username:w5"}
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userid:%s, username:%s", user.id, user.name) //这个不用加\n
				userInfos = append(userInfos, userInfo)
			}

			//最终写到管道中，一定是一个字符串
			r := strings.Join(userInfos, "\n") //连接数字切片，生成字符串
			//strings.Split() //分割字符串
			//`
			//	"userid:z3, username:z3"
			//	"userid:l4, username:l4"
			//	"userid:w5, username:w5"
			//`

			//将数据返回给查询的客户端
			newUser.msg <- r
			// netcat 127.0.0.1 8080 输入 \rename|xxx
		} else if len(userInput) > 9 && userInput[:7] == "\\rename" {
			//[:3]  // 0, 1, 2  ==> 左闭右开

			//	规则：  rename|Duke
			//1. 读取数据判断长度7，判断字符是rename
			//2. 使用|进行分割，获取|后面的部分，作为名字
			//func Split(s, sep string) []string
			//arry := strings.Split(userInput, "|")
			//name := arry[1]
			//3. 更新用户名字newUser.name = Duke
			newUser.name = strings.Split(userInput, "|")[1]
			allUsers[newUser.id] = newUser //更新map中的user

			//4. 通知客户端，更新成功
			newUser.msg <- "rename successfully!"
		} else {
			//如果用户输入的不是命令，只是普通的聊天信息，那么只需要写到广播通道中即可，由其他的go程进行常规转发
			message <- userInput
		}

		restTimer <- true
		//-------- 业务逻辑处理  结束----------
	}
}

//向所有的用户广播消息,启动一个全局唯一go程
func broadcast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast 程序退出!")

	for {
		//1. 从message中读取数据
		fmt.Println("broadcast监听message中...")
		info := <-message

		fmt.Println("message 接收到消息:", info)

		//2. 将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			//如果msg是非缓冲的，那么会在这里阻塞
			user.msg <- info
		}
	}
}

//每个用户应该还有一个用来监听自己msg管道的go程，负责将数据返回给客户端
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("111111 user : %s 的go程正在监听自己的msg管道:\n", user.name)
	for data := range user.msg {
		fmt.Printf("user : %s 写回给客户端的数据为:%s\n", user.name, data)

		//Write(b []byte) (n int, err error)
		_, _ = conn.Write([]byte(data))
	}
}

//i,j int
//启动一个go程，负责监听退出信号，触发后，进行清零工作: delete map, close conn都在这里处理
func watch(user *User, conn net.Conn, isQuit, restTimer <-chan bool) {
	fmt.Println("222222 启动监听退出信号的go程....")
	defer fmt.Println("watch go程退出!")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already!\n", user.name)
			fmt.Println("删除当前用户:", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-time.After(5 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit already!\n", user.name)
			fmt.Println("删除当前用户:", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-restTimer:
			fmt.Printf("连接%s 重置计数器!\n", user.name)
		}
	}
}
