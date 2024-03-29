package main

import (
	"wovert/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()


	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("[" , conn , "]收到client发来的数据：", msg)
	}
}

func main() {

	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	// 关闭监听
	defer listen.Close()

	fmt.Println("服务器等待客户端建立连接")
	for {
		// 阻塞等待客户端连接请求
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("服务器与客户端成功建立连接")
		go process(conn)
	}
}