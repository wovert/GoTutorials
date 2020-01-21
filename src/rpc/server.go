package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello Panda")
}

type Panda int

/**
 * 函数关键字（对象）函数名(对端发送过来的内容, 返回给对端的内容) 错误返回值
 */
func (this *Panda)Getinfo(argType int, replyType *int) error {
	fmt.Println("打印对端发送过来的内容为：", argType)

	// 修改内容值
	*replyType = argType + 88
	return nil
}

func main() {
	// 页面请求
	http.HandleFunc("/panda", pandatext)


	pd := new(Panda) // 将类实例话对象
	rpc.Register(pd) // 服务端注册一个对象
	rpc.HandleHTTP() // 连接到HTTP

	// 监听 http服务
	handle, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(handle, nil)
}
