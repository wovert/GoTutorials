package main

import (
  "fmt"
  "net"
  "os"
)

func errFunc2(err error, info string) {
  if err != nil {
    fmt.Println(info, err)
    // return // 返回当前函数调用
    // runtime.Goexit() // 结束当前go程
    os.Exit(1) // 结束当前进程
  }
}

func main() {
  host := "127.0.0.1:8080"
  conn, err := net.Dial("tcp", host)
  errFunc2(err, "Dial")
  defer conn.Close()

  httpRequest := "GET /helloworld HTTP/1.1\r\nHost:127.0.0.1:8080\r\n\r\n"
  conn.Write([]byte(httpRequest))

  buf := make([]byte, 4096)
  n, _ := conn.Read(buf)
  if n == 0 {
    return
  }

  fmt.Printf("|%s|\n", string(buf[:n]))

}