package main

import "fmt"

func slice(){
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6] // [2,3,4,5]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
}