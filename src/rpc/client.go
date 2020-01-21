package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络连接失败")
	}

	var pd int
	err = cli.Call("Panda.Getinfo", 1000, &pd)
	if err != nil {
		fmt.Println("远程过程调用失败")
	}
	fmt.Println("调用结果是：", pd)
}
