package main

import (
	"net"
	"fmt"
	"time"
)

func main() {

	// 组织一个 udp 地址结构, 指定服务器的IP+port
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("udp 服务器地址结构，创建完程!!!")
	// 创建用户通信的 socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp 服务器通信socket创建完成!!!")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)

	for {
		// 返回3个值，分别是 读取到的字节数， 客户端的地址， error
		n, cltAddr, err := udpConn.ReadFromUDP(buf)         // --- 主go程读取客户端发送数据
		if err != nil {
			fmt.Println("ReadFromUDP err:", err)
			return
		}
		// 模拟处理数据
		fmt.Printf("服务器读到 %v 的数据：%s\n", cltAddr, string(buf[:n]))

		go func() {                 // 每有一个客户端连接上来，启动一个go程 写数据。
			// 提取系统当前时间
			daytime := time.Now().String() + "\n"

			// 回写数据给客户端
			_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
			if err != nil {
				fmt.Println("WriteToUDP err:", err)
				return
			}
		}()
	}
}