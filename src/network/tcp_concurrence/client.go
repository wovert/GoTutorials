package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 服务器IP地址和端口，创建通信套接字
	conn, err := net.Dial("tcp", "192.168.3.12:8989")
	if err != nil {
		fmt.Println("net.Dial err:", err)
	}
	defer conn.Close()

	// 获取用户键盘输入（stdin），将输入数据发送给服务器
	go func(){
		for {
			str := make([]byte, 4096)
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			// 写给服务器，读多少，写多少
			conn.Write(str[:n])

		}
	}()

	// 回显服务器回发的大些数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		if n == 0 {
			fmt.Println("检测到服务器已经关闭")
			return
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("客户端读到服务器回发:", string(buf[:n]))
	}





}
