package main

import (
	"fmt"
	"net"
	"os"
)

func errInfo(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		//return
		//runtime.Goexit() // 结束当前协程
		os.Exit(1) // 将当前进程结束
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errInfo(err, "Dial:")
	defer conn.Close()

	// 模拟浏览器
	httpRequest := "GET /index HTTP/1.1\r\n"
	httpRequest += "Host: 127.0.0.1:8000\r\n"
	httpRequest += "\r\n"

	conn.Write([]byte(httpRequest))

	// 读取服务端返回的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errInfo(err, "没有读取到数据：")
	fmt.Printf("|%s\n|", string(buf[:n]))
}