package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4044")

	if err != nil {
		panic(err)
	}

	fmt.Println("listen to 4044")
	for {
		// 监听到新的连接，创建新的 goroutine 交给 handleConn函数 处理
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("新连接：", conn.RemoteAddr())

	result := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			fmt.Println("recv:", result.String())
		}
		result.Reset()
	}
}