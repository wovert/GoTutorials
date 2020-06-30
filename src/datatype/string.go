package main

import "fmt"

func main() {
	str := "hello" +
			" world!"
	fmt.Println(str)

	// 字符串都隐藏了一个结束符 ‘\0’
}