package tcp_upload

import (
	"fmt"
	"net"
	"strings"
)

// 处理用户请求
func HandleConn2(conn net.Conn) {
	// 函数调用完毕，自动关闭conn
	defer conn.Close()

	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Printf("客户端[%s]连接成功.\n", addr)

	buf := make([]byte, 2048)

	for {
		// 读取请求数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取客户端发送的数据错误: ", err)
			return
		}
		fmt.Printf("客户端[%s]发送的数据是: %s.\n", addr, string(buf[:n-1]))
		fmt.Printf("len = %d\n", len(string(buf[:n-1])))
		// fmt.Printf("len = %d\n", len(string(buf[:n-1])))
	
		// 客户端主动退出，Windows结束符\r\n, Linux结束符 \n
		if "exit" == string(buf[:n-1]) {
			fmt.Printf("客户端[%s]退出.\n", addr)
			conn.Close() // 客户端退出
			return
		}

		// 数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func Tcp2() {

	networkType := "tcp"
	host := "127.0.0.1:8000"

	// 1. 监听
	listener, err := net.Listen(networkType, host)
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	// 5. 关闭服务
	defer listener.Close()

	// 2. 阻塞等待用户请求
	// 接受多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			// return
			continue
		}
		

		// 处理用户请求，新建一个协程
		go HandleConn2(conn)

		// 3. 接受用户的请求数据
		// buf := make([]byte, 1024) // 1024byte的缓冲区
		// n, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("receive data error: ", err)
		// 	// return
		// 	continue
		// }
		// fmt.Println("buffer = ", string(buf[:n]))

		// defer conn.Close() // 关闭当前用户连接

	}
	// 4. 发送数据给用户
}