package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占用的字节数为 %d ", n2, unsafe.Sizeof(n2))
}