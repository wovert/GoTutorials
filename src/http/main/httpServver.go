package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Header:", r.Header)
  fmt.Println("URL:", r.URL)
  fmt.Println("Host:", r.Host)
  fmt.Println("Method:", r.Method)
  fmt.Println("RemoteAddr:", r.RemoteAddr)
  fmt.Println("Body:", r.Body)

  w.Write([]byte("hello world")) // 响应给客户端
}

func main() {

  // 1.监听地址和端口
  host := "127.0.0.1:8080"

  // 2.注册回调函数，该回调函数在服务器被访问时，自动被调用
  // 2.1 用户访问位置
  // 2.2 回调函数——必须是(w http.RequestWriter, r *http.Request)作为参数
  // http.RequestWriter 会写给客户端数据
  // *http.Request 请求参数数据
  http.HandleFunc("/helloworld", handler)

  // 3.绑定服务器监听地址和端口
  http.ListenAndServe(host, nil)

}