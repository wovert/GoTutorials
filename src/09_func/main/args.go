package main


import (
	"fmt"
)

func test(str string, args ...int) {
	println("str=", str)
	for index, value := range args {
		fmt.Printf("%d=%d\n", index, value)
	}
}

func main() {
	test("abc", 20, 30, 40)
}