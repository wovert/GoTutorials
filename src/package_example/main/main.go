package main

import (
	"package_example/calc"
	p "package_example/person"
	_ "package_example/test" // 不适应此包时使用 _ 字符
	"fmt"
	"time"
)

const (
	Man = 1
	Female = 2
)

func main() {
	sum := calc.Add(20, 30)
	sub := calc.Sub(100,20)
	// p.Test()
	name := p.Name
	age := p.Age
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
	fmt.Println("Name is ", name)
	fmt.Println("Age is ", age)


	for {
		second := time.Now().Unix()
		if (second % Female == 0) {
			fmt.Println("female")
		} else {
			fmt.Println("Male")
		}
		// Millisecond 毫秒
		// Microsecond 微妙
		time.Sleep(1000 * time.Millisecond)
	}

}
// 编译: go build package_example/main
// 指定编译目录：go build -o bin/my_calc.exe package_example/main
