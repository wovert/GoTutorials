package main

import (
	"workspace/messages" // 包的路径
	"workspace/users"
)

func main() {
	test()
	demo()
	message.Send() // 包名.函数()
	user.Login() // 包名.函数()
}