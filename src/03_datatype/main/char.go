package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var c1 byte // 声明字符类型
	c1 = 97
	c2 := 'B'
	fmt.Printf("c1 的类型 %T c1占用的字节数为 %d 字符：%c, 整数:%d\n ", c1, unsafe.Sizeof(c1), c1, c1)
	fmt.Printf("c2 的类型 %T c2占用的字节数为 %d 字符：%c, 整数:%d\n ", c2, unsafe.Sizeof(c2), c2, c2)
}