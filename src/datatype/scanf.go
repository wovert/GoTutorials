package main


import "fmt"

func main() {
	var a,b int

	// 阻塞等待用户输入
	fmt.Scanf("%d", &a)
	fmt.Scan(&b)
	fmt.Println("a=",a,"b=",b)

}