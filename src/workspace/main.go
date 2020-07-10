package main

import (
	"workspace/messages" // 包的路径
	"workspace/users"
	_ "workspace/demo"
	. "workspace/wovert"
)

func main() {
	// test()
	// demo()
	message.Send() // 包名.函数()
	user.Login() // 包名.函数()
	Run()
}