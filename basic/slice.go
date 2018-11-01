package main

import (
	"fmt"
)

func main() {
	x := [10]int{1,2,3,4,5,6,7,8,9,10}

	y := x[1:3]
	x[1] = 1000

	// 10 长度3，容量5
	s := make([]int, 3, 5)
	s = append(s, 5, 6, 7)

	fmt.Printf("%v", y) // [1000 3]
	fmt.Printf("%v", s) // [0 0 0 5 6 7]
	fmt.Printf("%v", cap(s)) // 自动扩充至 10
}
