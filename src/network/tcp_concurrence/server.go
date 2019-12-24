package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn){
	defer conn.Close()

	// 获取连接的客户端 Addr
	addr := conn.RemoteAddr()
	fmt.Println("客户端成功连接：", addr)


	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		fmt.Println(buf[:n]) // 输出 ACSII 13:回车符 \r, 10:换行符 \n
		if "exit\r\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务端接收到客户端退出请求，服务器关闭")
			return
		}

		if n == 0 {
			fmt.Println("服务端监测到客户端已经关闭，断开连接！")
			return
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		// 小写转大些
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

		// 处理数据-打印
		fmt.Println("服务器读到客户端数据：", string(buf[:n]))
	}
}

func main() {
	// 服务器通信协议、IP地址和端口号；创建一个用于监听的 socket
	listener, err := net.Listen("tcp", "192.168.3.12:8989")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 阻塞监听客户端连接请求，成功连接连接，返回用于通信的socket
	for {
		fmt.Println("服务器等待客户端建立连接...")
		conn, err:= listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() err:", err)
			return
		}

		// 服务器和客户端数据通信
		go HandlerConnect(conn)
	}

}
