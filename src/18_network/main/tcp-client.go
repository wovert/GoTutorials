package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("连接服务器错误：", err)
		return
	}

	// main 调用完毕，关闭连接
	defer conn.Close()

	go func() {
		// 从键盘输入内容，给服务器发送数据
		str := make([]byte, 1024)
		for {
			// 从键盘读取内容，放在str
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("从键盘读取内容错误：", err)
				return
			}
			// 把输入的内容给服务器发送
			conn.Write(str[:n])
		}
	}()

	// 接受服务器回复数据，新任务
	// 切片缓存
	buf := make([]byte, 1024)
	// 不停的接受数据
	for {
		// 接受服务器的请求
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接受服务器数据错误：", err)
			return
		}
		// 打印接受到的内容
		fmt.Printf("接受服务端数据：%s.\n", string(buf[:n-1]))
	}

	// conn.Write([]byte("Are u OK?"))


}