package main

import (
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	// w：写回给客户端
	// r: 从客户端浏览器读到的数据
	w.Write([]byte("Hello World"))
}

func main() {

	// 注册回调函数，回调函数在服务器被访问时，自动被调用
	http.HandleFunc("/index", handler)

	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
