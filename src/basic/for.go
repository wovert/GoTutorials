package main

import "fmt"

func main() {
	//var x int
	//sum := 0
	//for x=1; x<=10; x++ {
	//	sum += x
	//}
	//fmt.Print(sum)
	//
	//for x < 100 {
	//	fmt.Print(x)
	//	x++
	//}

	// 死循环
	//for {
	//	// code...
	//}

	p := make(map[string]int)
	p["zhansan"] = 10
	p["lisi"] = 30
	p["wangwu"] = 20

	for k,v := range p {
		fmt.Println(k, v)
	}

	name := "wovert"
	for k,v := range name {
		fmt.Printf("%d=%c\n", k,v) // index ASCII
	}
}