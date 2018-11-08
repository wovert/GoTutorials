package main

import (
	"fmt"
)

func main() {
	//var student map[string]float32
	//student = make(map[string]float32)

	// 简写形式
	student := make(map[string]float32)

	student["name"] = 19.2
	fmt.Printf("%v", student)


}
