package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务器通信协议、IP地址和端口号；创建一个用于监听的 socket
	listener, err := net.Listen("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close() // 关闭监听

	fmt.Println("服务器等待客户端建立连接...")

	// 阻塞监听客户端连接请求，成功连接连接，返回用于通信的socket
	conn, err:= listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务器与客户端成功连接")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	// 读多少写多少
	conn.Write(buf[:n])

	// 处理数据-打印
	fmt.Println("服务器读到数据：", string(buf[:n]))
}
