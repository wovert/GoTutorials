package main

import (
	"fmt"
)

func main() {
	type bigint int64
	var a bigint
	a = 200
	type (
		long int64
		char byte
	)
	var b long = 11
	var ch char = 'a'
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("ch=", ch)
}