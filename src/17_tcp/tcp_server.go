package main

import (
	// "wovert/proto"
	// "bufio"
	"fmt"
	// "io"
	"net"
)

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()

	// 接受用户的请求 1024大小的缓冲区
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf=", string(buf[:n]))
}

func main() {

	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("监听端口错误:", err)
		return
	}

	// 关闭监听
	defer listen.Close()

	fmt.Println("服务器[等待]客户端建立连接")
	for {
		// 阻塞等待客户端连接请求
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("服务器与客户端[成功建立]连接")
		go process(conn)
	}
}