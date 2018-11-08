package main

// 查找 GOPATH 环境变量的目录中 test 包
import "test"

func main() {
	test.Hello()
}