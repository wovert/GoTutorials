package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 创建用户结构体类型
type Client struct {
	C chan string
	Name string
	Addr string
}

/**
 * 定义全局变量
 */

var onlineMap map[string]Client // 创建全局map在线用户, key:string, value:Client
var message = make(chan string) // 创建全局 channel，用于传递用户消息

/**
 * 错误输出信息
 */
func outputError(msg string) {
	fmt.Println("msg")
}

/**
 * 管理全局消息通道和各个客户端消息通道
 */
func Manager() {

	onlineMap = make(map[string]Client) // 初始化 onlineMap
	for { // 监听全局channel是否有数据
		msg := <- message // 有数据存储到 msg，无数据阻塞
		for _, client := range onlineMap { // 循环发送消息给所有在线用户
			client.C <- msg
		}
	}
}

/**
 * 响应给客户端消息
 */
func WriteMsgToClient(client Client, conn net.Conn) {
	for msg := range client.C { 	// 监听用户自带的channel上是否有消息
		conn.Write([]byte(msg + "\n")) 		// 发送消息给当前客户端
	}
}

/**
 * 生成发送消息内容
 */
func MakeMsg(client Client, msg string)(buf string) {
	buf = "[" + client.Addr + "]" + client.Name + ": " + msg
	return
}


/**
 * 接受客户端请求连接
 */
func HandlerConnect(conn net.Conn) {
	defer conn.Close() 	// 关闭客户读连接请求

	isOnline := make(chan bool) // 创建channel判断，用户是否活跃
	netAddr := conn.RemoteAddr().String() // 获取用户网络地址 IP+PORT

	client := Client { // 创建新连接用户的结构体，默认用户名是 IP+PORT
		make(chan string),
		netAddr,
		netAddr,
	}

	onlineMap[netAddr] = client // 新连接用户添加到在线用户 map 中，key=IP+PORT, value:client

	go WriteMsgToClient(client, conn) // 创建专门用户给当前用户发送消息的协程，不会随着HandlerConnect go程结束了结束WriteMsgToClient go程

	message <- MakeMsg(client, "login")  	// 发送用户上线消息到全局 channel 中

	isQuit := make(chan bool) 	// 创建一个channel，判断退出状态

	go func() { // 创建匿名子协程，处理用户发送的消息
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true // 没有读到消息内容，发送退出管道
				fmt.Printf("监测到客户端: %s退出\n", client.Name)
				return
			}
			if err != nil {
				fmt.Println("读取错误：", err)
				return
			}

			msg := string(buf[:n]) // 读到的用户消息，保存到msg中

			if (msg == "who\n" && len(msg) == 4) || (msg == "who\r\n" && len(msg) == 5) { // 提取在线用户列表
				conn.Write([]byte("online  user list: \n"))

				for _, user := range onlineMap { // 遍历map在线用户
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" { // rename|
				newName := strings.Split(msg, "|")[1]
				fmt.Println("new Name:", newName)
				client.Name = newName // 修改结构体
				onlineMap[netAddr] = client // 更新用户列表 onlineMap
				conn.Write([]byte("rename sucessfull.\n"))
			} else {
				message <- MakeMsg(client, msg) // 将读到的用户消息，写入到message中进行广播
			}

			isOnline <- true // 活跃用户

			//fmt.Println([]byte(msg))
			//fmt.Println(msg + "...")
		}
	}()

	// 保证不退出
	for {
		select { // 监听channel上的流动
			case <-isQuit: // 读取isQuit值
				close(client.C) // 关闭客户端 监听用户自带的channel上是否有消息， WriteMsgToClient go程的循环结束，即Go程退出
				delete(onlineMap, client.Addr) // 将用户从online移除
				message <- MakeMsg(client, "logout") // 写入用户退出消息到全局channel
				return
			case <-isOnline: // 是否有消息通信，有发送内容就重置下面的计时器
			case <-time.After(time.Second * 10): // 超出10s剔除
				delete(onlineMap, client.Addr) // 将用户从online移除
				message <- MakeMsg(client, "logout") // 写入用户退出消息到全局channel
				return
		}
	}
}

func main() {
	// 定义服务器地址和端口
	serverAddr := "127.0.0.1"
	serverPort := "9898"
	netType := "tcp"

	serverAddrPort := serverAddr + ":" + serverPort
	listener, err := net.Listen(netType, serverAddrPort) // 1. 创建监听套接字
	if err != nil {
		outputError("创建监听套接字错误")
		return
	}

	defer listener.Close() // 2. 关闭套接字

	go Manager() // 3. 创建管理者协程，管理map和全局channel

	for { // 4. 循环监听客户端连接请求
		conn, err := listener.Accept()
		if err != nil {
			outputError("监听客户端连接错误")
			return
		}
		go HandlerConnect(conn) // 4.x 启动协程处理客户端数据请求
	}
}