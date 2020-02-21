package main

import (
	pd "../myproto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 定义类
type Server struct {}

// 第一个接口：打招呼的服务
// context，客户端发送过来的参数，返回给客户端的参数，错误
func (this *Server)SayHello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloRes, err error) {
	return &pd.HelloRes{Msg: "hello " + in.Name}, nil
}
// 第二个接口：说名字的服务
func (this *Server)SayName(ctx context.Context, in *pd.NameReq) (out *pd.NameRes, err error) {
	return &pd.NameRes{Msg: in.Name + " 早上好"}, nil
}

func main() {
	// 创建网络
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误：", err)
	}

	srv := grpc.NewServer() 	// 创建grpc服务
	pd.RegisterHelloServerServer(srv, &Server{}) // 注册服务


	// 等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误：", err)
	}
}