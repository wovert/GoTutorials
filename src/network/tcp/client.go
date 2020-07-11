package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务器IP地址和端口，创建通信套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Println("net.Dial err:", err)
	}
	defer conn.Close()

	// 主动协数据给服务器
	conn.Write([]byte("Hi中国"))

	// 接受服务器响应数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Printf("服务器响应:%s, 字节数：%d\n", string(buf[:n]), n)
}
