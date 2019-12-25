package main

import (
	"fmt"
	"net"
)

// 创建用户结构体类型
type Client struct {
	C chan string
	Name string
	Addr string
}

// 创建全局map在线用户, key:string, value:Client
var onlineMap map[string]Client


// 创建全局 channel，用于传递用户消息
var message = make(chan string)

// 错误输出信息
func outputError(msg string) {
	fmt.Println("msg")
}

func Manager() {
	// 初始化 onlineMap
	onlineMap = make(map[string]Client)

	// 监听全局channel是否有数据
	for {
		// 有数据存储到 msg，无数据阻塞
		msg := <- message

		// 循环发送消息给所有在线用户
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

func WriteMsgToClient(client Client, conn net.Conn) {
	// 监听用户自带的channel上是否有消息
	for msg := range client.C {
		// 发送消息给当前客户端
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect(conn net.Conn) {
	// 关闭客户读连接请求
	defer conn.Close()

	// 获取用户网络地址 IP+PORT
	netAddr := conn.RemoteAddr().String()

	// 创建新连接用户的结构体，默认用户名是 IP+PORT
	client := Client {
		make(chan string),
		netAddr,
		netAddr,
	}

	// 新连接用户添加到在线用户 map 中，key=IP+PORT, value:client
	onlineMap[netAddr] = client

	// 创建专门用户给当前用户发送消息的协程
	go WriteMsgToClient(client, conn)

	// 发送用户上线消息到全局 channel 中
	message <- "[" + netAddr + "]" + client.Name + " login"

	// 保证不退出
	for {
		;
	}
}

func main() {
	server_addr_port := "192.168.3.12:9898"

	// 1. 创建监听套接字
	listener, err := net.Listen("tcp", server_addr_port)
	if err != nil {
		outputError("创建监听套接字错误")
		return
	}

	// 2. 关闭套接字
	defer listener.Close()

	// 3. 创建管理者协程，管理map和全局channel
	go Manager()

	// 4. 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			outputError("监听客户端连接错误")
			return
		}

		// 4. 启动协程处理客户端数据请求
		go HandlerConnect(conn)
	}



}
