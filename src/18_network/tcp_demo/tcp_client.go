package main

import (
	// "wovert/proto"
	"fmt"
	"net"
	"strconv"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	for i := 0; i < 20; i++ {
		str := strconv.Itoa(i)
		msg := "Hello, Hello. How are you?" + str + "\n"
		conn.Write([]byte(msg))
	}
}