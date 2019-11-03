package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var f1 float64
	f1 = 3.14
	fmt.Printf("f1 的类型 %T f1占用的字节数为 %d ", f1, unsafe.Sizeof(f1))
}