package main


import (
	"fmt"
	"os"
)

func test(str string, args ...int) {
	println("str=", str)
	for index, value := range args {
		fmt.Printf("%d=%d\n", index, value)
	}
}

func main() {
	test("abc", 20, 30, 40)

	// 接受用户的参数，字符串方式传递
	list := os.Args
	n := len(list)
	fmt.Println("n=", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d=%s\n", i, list[i])
	}
}