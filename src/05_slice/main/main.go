package main

import (
	"fmt"
)

func slice1() {
	arr := [6]int {1,2,3,4,5,6}
	s := arr[1:3:5]

	fmt.Println("s=", s) // [2 3]
	fmt.Println("len(s)=", len(s)) // 2
	fmt.Println("cap(s)", cap(s)) // 4
}

func slice2() {
	arr := [8]int {1,2,3,4,5,6,7,8}
	s := arr[1:5]

	fmt.Println("s=", s) // [2 3 4 5]
	fmt.Println("len(s)=", len(s)) // 4
	fmt.Println("cap(s)", cap(s)) // 7
}
func main() {
	slice1()
	slice2()
}