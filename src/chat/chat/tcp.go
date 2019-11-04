package chat

import (
	"fmt"
	"net"
	"time"
	"strings"
)

// 客户端信息
type Client struct {
  C chan string	// 用户发送数据的管道
  Name string	// 用户名
  Addr string	// 网络地址
}

// 在线用户
var onlineMap map[string]Client

// 用户发送消息管道
var message = make(chan string)

func makeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func maneger() {
	// 给 map 分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message // 没有消息前，这里会阻塞
		
		// 遍历onlineMap，给每个成员发送此消息
		for _, cli := range onlineMap {
		  cli.C <- msg
		}
	}
}

// 发送信息给客户端
func writeMsgToClient(cli Client, conn net.Conn) {
  for msg := range cli.C {
	// 给当前客户端发送信息
	conn.Write([]byte(msg + "\n"))
  }
}

// 用户处理连接
func handleConn(conn net.Conn) {
  // 获取客户端的网络地址
  cliAddr := conn.RemoteAddr().String()

  // 创建一个结构体，默认用户名和网络地址一样
  cli := Client{make(chan string), cliAddr, cliAddr}

  // 把结构体添加到 map
  onlineMap[cliAddr] = cli

  // 新开一个协程，专门给当前客户端发送信息
  go writeMsgToClient(cli, conn)


  // 广播某个人在线
  message <- makeMsg(cli, "login")

	// 提示我是谁
	cli.C <- makeMsg(cli, "我在线")

	isQuit := make(chan bool) // 对方是否主动退出
	hasData := make(chan bool) // 对方是否数据发送

  // 新建一个协程，接受用户发送过来的数据
  go func(){
		buf := make([]byte, 2048)
		
		// 循环读取发送的内容
		for {
			n, err := conn.Read(buf)
			fmt.Println("n=", n)
			if n == 0 || n == 1 { // 对方断开 or 出问题
				isQuit <- true
				// fmt.Println(cli.Name + "退出")
				fmt.Println("conn.Read err = ", err)
				return
			}
			
			// fmt.Println("start:" + string(buf[:n]) +"====")
			// fmt.Println("len:", len(string(buf[:n])))
			if (len(string(buf[:n])) == 1) {
				// return
			}
			msg := string(buf[:n-2]) // 通过 windows nc测试，多一个换行
			// fmt.Println(len(msg))

			// 查询在线用户有多少人
			if len(msg) == 3 && msg == "who" {
				// 遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list :\n"))

				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				// rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
        // 转发此内容
        message <- makeMsg(cli, string(msg))
			}
			hasData <- true // 有数据
		}
  }()

  for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr) // 当前用户从map移除
			message <- makeMsg(cli, "logout") // 告诉所有在线用户，谁退出了
			return
		case <- hasData:
		case <- time.After(30 * time.Second): // 30s后退出聊天室
		  delete(onlineMap, cliAddr) // 当前用户从map移除
		  message <- makeMsg(cli, "time out leave out") // 告诉所有在线用户，谁退出了
		  return
		}
  }
}

// 聊天入口
func TcpEntry () {
  listener, err := net.Listen("tcp", ":8000")
  if err != nil {
	fmt.Println("net.Listen err = ", err)
	return
  }

  defer listener.Close()
  
  // 新开一个协程，转发消息，只要有消息来了，遍历onlineMap, 给map每个成员发送消息
  go maneger()

  // 主协程, 循环阻塞等待用户连接
  for {
	  conn, err := listener.Accept()
	  if err != nil {
		  fmt.Println("listener.Accept err = ", err)
		  continue
	  }
	  go handleConn(conn) // 处理用户连接
  }
}