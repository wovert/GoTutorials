package main

import (
	"fmt"
)
func slice0() {
	intArr := [...]int{1,2,3,4,5,6}
	slice := intArr[0:3:4] // index:1-3(不包含3)

	fmt.Printf("intArr[0]=%v\n", &intArr[0])
	fmt.Printf("slice[0]=%v\n", &slice[0])
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("len(slice)=%v\n", len(slice))
	fmt.Printf("cap(slice)=%v\n", cap(slice))
	fmt.Printf("intArr=%v\n", intArr)
	slice[1] = 34
	fmt.Printf("intArr=%v\n", intArr)
}

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
	slice0()
	// slice1()
	// slice2()
}