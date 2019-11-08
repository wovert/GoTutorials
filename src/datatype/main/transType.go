package main

import "fmt"

func main() {
	var i int32 = 257
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	var n4 byte = byte(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v n4=%v\n", i, n1, n2, n3, n4)
	fmt.Printf("i=%T \n", i)
	fmt.Println("==================")
	demo()
	fmt.Println("==================")
	demo2()
}

func demo () {
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println("n2=", n2, "n3=", n3)
}

func demo2() {
	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 编译通过，但是溢出
	n3 = int8(n1) + 128 // 编译不通过（int8:127~-128）
	fmt.Println(n3, n4)
}