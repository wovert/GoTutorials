package main

import (
	"package_example/calc"
	"fmt"
)

func main() {
	sum := calc.Add(20, 30)
	sub := calc.Sub(100,20)
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}
// 编译: go build package_example/main
// 指定编译目录：go build -o bin/my_calc.exe package_example/main
