package main

import (
	pd "../myproto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常:", err)
	}
	defer conn.Close() // 网络延迟关闭


	c := pd.NewHelloServerClient(conn) // 获得grpc 句柄
	res, err := c.SayHello(context.Background(), &pd.HelloReq{Name: "微明"}) // 通过句柄调用函数
	if err != nil {
		fmt.Println("SayHello 服务调用失败")
	}
	fmt.Println("调用 SayHello 的返回: ", res.Msg)

	re, err := c.SayName(context.Background(), &pd.NameReq{Name: "CR7"})
	if err != nil {
		fmt.Println("SayName 服务调用失败")
	}
	fmt.Println("调用 SayName 的返回: ", re.Msg)

}