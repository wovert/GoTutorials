package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
func main(){
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6] // [2,3,4,5]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	// 声明 slice, Zeor vlaue for slice
	// s == nil
	var s1 []int
	for i := 0; i < 100; i++ {
		printSlice(s1)
		s1 = append(s1, 2 * i + 1)
	}
	fmt.Println(s1)

	// create slice method
	s2 := []int{2,4,6,7,8,10,29}
	printSlice(s2)

	// 16:len, 20:cap
	s3 := make([]int, 16)

	// 32: cap
	s4 := make([]int, 10, 32)

	printSlice(s3)
	printSlice(s4)

	fmt.Println("=========== Copy slice ===========")
	copy(s3, s2) // s3复制到s2
	printSlice(s3) // [2 4 6 7 0 0 0 0 0....]

	fmt.Println("=========== Delete elemes from slice ===========")
	s3 = append(s3[:3], s3[4:]...) // delete 7
	printSlice(s3) // [2 4 6 0 0 0 0 0....]

	fmt.Println("=========== Pop from front ===========")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2) // [4 6 0 0 0 0 0....]
	fmt.Println(front)

	fmt.Println("=========== Pop from back ===========")
	tail := s2[len(s2)-1]
	s2 = s3[:len(s2)-1]
	printSlice(s2) // [4 6 0 0 0 0 0....]
	fmt.Println(tail)

}