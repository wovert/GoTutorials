package main 

import ("fmt")

var g = 10
// g2 := 100 // 不能用在函数体外部

func main(){
	// 定义变量
	var x int
	x = 100
	y := "wovert"
	fmt.Printf("%d %s %d %d", x, y, g)
}