package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 创建用户结构体类型
type Client struct {
	C    chan string // 通知连接用户消息
	Name string
	Addr string
}

/**
 * 定义全局变量
 */

// 创建全局map在线用户, key:string, value:Client
var onlineMap map[string]Client

// 创建全局 channel，用于传递用户消息
var message = make(chan string)

/**
 * 错误输出信息
 */
func outputError(msg string) {
	fmt.Println("msg")
}

/**
 * 管理全局消息通道和各个客户端消息通道
 */
func boardCast() {

	// 广播消息
	onlineMap = make(map[string]Client) // 初始化 onlineMap
	for {                               // 监听全局channel是否有数据
		msg := <-message                   // 有数据存储到 msg，无数据阻塞
		for _, client := range onlineMap { // 循环发送消息给所有在线用户
			client.C <- msg
		}
	}
}

/**
 * 响应给客户端消息
 */
func WriteMsgToClient(client *Client, conn net.Conn) {
	// 监听用户自带的channel上是否有消息
	for msg := range client.C {
		// 发送消息给当前客户端
		conn.Write([]byte(msg + "\n"))
	}
}

/**
 * 生成发送消息内容
 */
func MakeMsg(client Client, msg string) (buf string) {
	buf = "[" + client.Addr + "]" + client.Name + ": " + msg
	return
}

/**
 * 接受客户端请求连接
 */
func HandlerConnect(conn net.Conn) {
	// 关闭客户读连接请求
	defer conn.Close()

	// 创建channel判断，用户是否活跃
	isOnline := make(chan bool)

	// 获取用户网络地址 IP+PORT
	netAddr := conn.RemoteAddr().String()

	// 创建新连接用户的结构体，默认用户名是 IP+PORT
	client := Client{
		make(chan string),
		netAddr,
		netAddr,
	}

	// 新连接用户添加到在线用户 map 中，key=IP+PORT, value:client
	onlineMap[netAddr] = client

	// 创建专门用户给当前用户发送消息的协程，不会随着HandlerConnect go程结束了结束WriteMsgToClient go程
	go WriteMsgToClient(&client, conn)

	// 发送用户上线消息到全局 channel 中
	message <- MakeMsg(client, "login")

	// 创建一个channel，判断退出状态
	isQuit := make(chan bool)

	// 处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			// n 读取数据长度
			n, err := conn.Read(buf)
			msg := string(buf[:n]) // 读到的用户消息，保存到msg中

			isLogout := (msg == "logout\n" && len(msg) == 7) || (msg == "who\r\n" && len(msg) == 8)
			if n == 0 || isLogout {
				// 没有读到消息内容，发送退出管道
				isQuit <- true
				fmt.Printf("监测到客户端: %s退出\n", client.Name)
				return
			}
			if err != nil {
				fmt.Println("读取错误：", err)
				return
			}

			if (msg == "who\n" && len(msg) == 4) || (msg == "who\r\n" && len(msg) == 5) { // 提取在线用户列表
				conn.Write([]byte("online  user list: \n"))

				for _, user := range onlineMap { // 遍历map在线用户
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
				// rename|
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				fmt.Println("New Name:", newName)

				// 修改结构体
				client.Name = newName

				// 更新用户列表 onlineMap
				onlineMap[netAddr] = client

				conn.Write([]byte("Rename sucessfull.\n"))
			} else {
				// 将读到的用户消息，写入到message中进行广播
				message <- MakeMsg(client, msg)
			}

			// 活跃用户
			isOnline <- true

			//fmt.Println([]byte(msg))
			//fmt.Println(msg + "...")
		}
	}()

	// 保证不退出
	for {
		// 监听channel上的流动
		select {
		case <-isQuit:
			// 关闭客户端 监听用户自带的channel上是否有消息， WriteMsgToClient go程的循环结束，即Go程退出
			close(client.C)
			delete(onlineMap, client.Addr)

			// 写入用户退出消息到全局channel
			message <- MakeMsg(client, "logout")
			return
		case <-isOnline: // 是否有消息通信，有发送内容就重置下面的计时器
		case <-time.After(time.Second * 30): // 超出30s剔除
			// 将用户从online移除
			delete(onlineMap, client.Addr)

			// 写入用户退出消息到全局channel
			message <- MakeMsg(client, "logout")
			return
		}
	}
}

func main() {
	// 定义服务器地址和端口
	serverAddr := "127.0.0.1"
	serverPort := "9898"
	netType := "tcp"

	serverAddrPort := fmt.Sprintf("%s:%d", serverAddr, serverPort)

	// 1. 创建监听套接字
	listener, err := net.Listen(netType, serverAddrPort)
	if err != nil {
		outputError("创建监听套接字错误")
		return
	}

	defer listener.Close() // 2. 关闭套接字

	go boardCast() // 3. 创建管理者协程，管理map和全局channel

	for { // 4. 循环监听客户端连接请求
		conn, err := listener.Accept()
		if err != nil {
			outputError("监听客户端连接错误")
			return
		}
		// 处理客户端数据请求
		go HandlerConnect(conn)
	}
}
