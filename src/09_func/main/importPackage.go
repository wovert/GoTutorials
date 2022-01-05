package main


import (
	. "fmt" // 调用函数，无序通过报名
	operatorSystem "os" // 给包起别名
	_ "wovert/09_func/other" // _ 表示仅调用包的init 函数
)

func main() {

	// 接受用户的参数，字符串方式传递
	list := operatorSystem.Args
	n := len(list)
	Println("n=", n)
	for i := 0; i < n; i++ {
		Printf("%d=%s\n", i, list[i])
	}
}