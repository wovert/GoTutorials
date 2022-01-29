package main

import (
	"net"
	"fmt"
	"os"
)

func recvFile(conn net.Conn, fileName string)  {
	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	// 从 网络中读数据，写入本地文件
	buf := make([]byte, 4096)
	for {
		n,_ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收文件完成。")
			return
		}
		// 写入本地文件，读多少，写多少。
		f.Write(buf[:n])
	}
}

func main()  {
	// 创建用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println(" net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(" listener.Accept() err:", err)
		return
	}
	defer conn.Close()

	// 获取文件名，保存
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(" conn.Read err:", err)
		return
	}
	fileName := string(buf[:n])

	// 回写 ok 给发送端
	conn.Write([]byte("ok"))

	// 获取文件内容
	recvFile(conn, fileName)
}