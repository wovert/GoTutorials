package main

import "fmt"

func slice(){
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6] // [2,3,4,5]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	fmt.Printf("arr[2]=%v\n", &arr[2])
	fmt.Printf("s[0]=%v\n", &s[0])
	// s[0] = 8
	arr[2] = 10
	fmt.Printf("arr[2]=%v, %v\n", &arr[2], arr[2])
	fmt.Printf("s[0]=%v, %v\n", &s[0], s[0])

}

func main() {
	slice()
}