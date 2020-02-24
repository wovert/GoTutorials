package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func slice0() {
	intArr := [...]int{1,2,3,4,5,6}
	slice := intArr[0:3:4] // index:1-3(不包含3)

	fmt.Printf("intArr[0]=%v\n", &intArr[0])
	fmt.Printf("slice[0]=%v\n", &slice[0])
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("len(slice)=%v\n", len(slice)) // 3-1=2
	fmt.Printf("cap(slice)=%v\n", cap(slice)) // 4-0=4
	fmt.Printf("intArr=%v\n", intArr)
	slice[1] = 34
	fmt.Printf("intArr=%v\n", intArr)
}

func slice1() {
	arr := [6]int {1,2,3,4,5,6}
	//s := arr[1:3:5]
	s := arr[1:3] // 6-1=5 cap=元素数量-开始解决索引值

	fmt.Println("s=", s) // [2 3]
	fmt.Println("len(s)=", len(s)) // 2
	fmt.Println("cap(s)", cap(s)) // 5
}

func slice2() {
	arr := [8]int {1,2,3,4,5,6,7,8}
	s := arr[1:5] // cap默认是8

	fmt.Println("s=", s) // [2 3 4 5]
	fmt.Println("len(s)=", len(s)) // 4
	fmt.Println("cap(s)", cap(s)) // 8-1=7
}

func slice3() {
	var slice []float64 = make([]float64, 3, 5)
	fmt.Printf("slice=%v\n", slice)
	slice[0] = 1.0
	slice[1] = 2.0
	slice[2] = 3.0
	// slice[3] = 4.0
	fmt.Printf("slice=%v\n", slice)
}

func slice4() {
	var slice []float64 = []float64 {1.0, 2.0, 3.0}
	fmt.Printf("slice=%v\n", slice)
	slice[2] = 10.0
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("len(slice)=%v\n", len(slice))
	fmt.Printf("cap(slice)=%v\n", cap(slice))
}

func slice5() {
	s1 := []int {1,2,4,5}
	s2 := append(s1, 6)
	s3 := append(s2, 9)
	s4 := append(s3, 10)
	s5 := append(s4, 11)
	s6 := append(s5, 12)
	s7 := append(s6, 15)
	s8 := append(s7, 16)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Println(s8)
}

func noEmpty(data []string) []string {
	out := data[:0] // 原切片上截取一个长度为0的切片 make([]string, 0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str // 直接在源串上操作
			i++
		}
	}
	return data[:i]
}

func strTest() {
	str := "hello@hotmail.com"
	slice := str[6:]
	fmt.Println("slice=", slice)

	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println("str=", str)
}

func fb(n int) ([]uint64) {
	fbSlice := make([]uint64, n)
	fbSlice[0] = 1
	fbSlice[1] = 1

	for i := 2; i < n; i++ {
		fbSlice[i] = fbSlice[i - 1] + fbSlice[i - 2]
	}
	return fbSlice
}

func main() {
	// slice0()
	//slice1()
	// slice2()
	// slice3()
	// slice4()
	//slice5()
	//return

	data := []string{"red", "", "black", "", "", "pink", "blue", "orange", "gray", "purple", "yellow"}
	afterData := noEmpty2(data)
	fmt.Printf("afterData:%v\n", afterData)
	return

	// s1 := afterData[6:] // [purple, yellow] 
	// s2 := afterData[0:5] // "red", black", pink", "blue", "orange"
	// copy(s2, s1)
	// fmt.Printf("s2=%v\n", s2) // [purple yellow pink blue orange]
	
	// strTest()

	// fbSlice := fb(10)
	// fmt.Println("fbSlice=", fbSlice)

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
	printSlice(s2) // [4 
	println(tail)
}