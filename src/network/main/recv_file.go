package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

// 接受文件内容
func RecvFile(fileName string, conn net.Conn) {

  // 新建文件
  f, err := os.Create(fileName)
  if err != nil {
	fmt.Println("os.Create err =", err)
	return
  }

  buf := make([]byte, 1024*4)

  // 接受多少，就写多少
  for {
	n, err1 := conn.Read(buf) // 读取发送的文件内容
	if err1 != nil {
	  if err1 == io.EOF {
	    fmt.Println("文件接受完毕")
	  } else {
		fmt.Println("conn.read err = ", err1)
	  }
	  return
	}

	if n == 0 {
	  fmt.Println("n == 0 文件接受完毕")
	  return
	}
	f.Write(buf[:n]) // 往文件写入内容
  }
}

func main() {
  litener, err := net.Listen("tcp", "127.0.0.1:8000")
  if err != nil {
	  fmt.Println("net.Listen err = ", err)
	  return
  }
  
  defer litener.Close()

  // 阻塞等待用户连接
  conn, err1 := litener.Accept()
  if err1 != nil {
    fmt.Println("litener.Accept err = ", err1)
    return
  }

  defer conn.Close()

  buf := make([]byte, 1024)
  var n int
  n, err = conn.Read(buf) // 读取客户端发送的文件名
  if err != nil {
    fmt.Println("conn.Read err = ", err1)
    return
  }

  fileName := string(buf[:n])

  // 回复"ok"
  conn.Write([]byte("ok"))

  // 接受文件内容
  RecvFile(fileName, conn)
}