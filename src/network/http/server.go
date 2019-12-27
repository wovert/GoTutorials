package main

import (
	"fmt"
	"net"
	"os"
	//"runtime"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		//return
		//runtime.Goexit() // 结束当前协程
		os.Exit(1) // 将当前进程结束
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8900")
	errFunc(err, "net.Listener err:")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "Accept err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read:")

	fmt.Printf("|%s|\n", string(buf[:n]))

}
