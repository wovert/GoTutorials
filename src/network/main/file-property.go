package main

import (
	"fmt"
	// "net"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage: xxx file")
		return
	}

	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	fmt.Println("name = ", info.Name())
	fmt.Println("size = ", info.Size()) // 字节
}