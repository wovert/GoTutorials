package main

import (
	"os"
	"fmt"
	"net"
	"io"
)

func sendFile(conn net.Conn, filePath string)  {
	// 只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()

	// 从本文件中，读数据，写给网络接收端。 读多少，写多少。原封不动。
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送文件完成。")
			} else {
				fmt.Println("os.Open err:", err)
			}
			return
		}
		// 写到网络socket中
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

func main()  {
	list := os.Args     // 获取命令行参数

	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件绝对路径")
		return
	}
	// 提取 文件的绝对路径
	filePath := list[1]

	//提取文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fileName := fileInfo.Name()

	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 发送文件名给 接收端
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	// 读取服务器回发的 OK
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	if "ok" == string(buf[:n]) {
		// 写文件内容给服务器——借助conn
		sendFile(conn, filePath)
	}
}