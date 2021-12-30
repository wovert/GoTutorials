package main

import "fmt"

func main() {
	str := "hello" +
		" world!"
	fmt.Println(str)

	// 字符串都隐藏了一个结束符 ‘\0’

	name := "我wovert"
	// 索引 切片 ASCII 因为 uint8
	msg := "abcdef"
	fmt.Printf("%T %#v %c\n", msg[0], msg[1], msg[2])
	fmt.Println(msg[1:3]) // 不包括3

	// len 字节大小
	fmt.Println(len(msg))
	fmt.Println(len(name))

	// 字数大小
	r := []rune(name)
	fmt.Println(len(r))
}
