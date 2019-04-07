package main

// 没有使用一定要注释
import (
"fmt"
// "time"
)

func add(a int, b int) int {
	var sum int
	sum =  a + b
	return sum
}

func main() {
	var c int 
	c = add(100, 200)
	// go test_goroute(300 ,300)
	fmt.Println(c)
	// fmt.Print("add(100,200)=", c)

	// for i := 0; i < 100; i++ {
	// 	// 子进程
	// 	go test_print(i)

	// }
	// // 主主进程
	// time.Sleep(10*time.Second)
	// test_pipe() // 3，有3个数据
}