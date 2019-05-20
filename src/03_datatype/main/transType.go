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
}