package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// udp 地址结构
	serverAddr, err := net.ResolveUDPAddr("udp", "192.168.3.12:9898")
	if err != nil {
		fmt.Println("ResolverUDPAddr err:", err)
		return
	}
	fmt.Println("UDP 服务器地址结构完成")

	// create udp socket
	udpConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("UDP 服务器通信SOCKET创建完成")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, clientAddr, err := udpConn.ReadFromUDP(buf) // 返回读取到的字节个数,客户端地址,错误
	if err != nil {
		fmt.Println("ReadFromUDP err:", err)
		return
	}
	fmt.Printf("服务器读到 %v的数据%s\n", clientAddr, string(buf[:n]))

	// 提取当前时间
	daytime := time.Now().String()
	// 回写数据给客户端
	_, err = udpConn.WriteToUDP([]byte(daytime), clientAddr)
	if err != nil {
		fmt.Println("WriteToUDP err:", err)
		return
	}

}
