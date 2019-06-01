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
	out := data[:0]
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func main() {
	// slice0()
	// slice1()
	// slice2()
	// slice3()
	// slice4()
	// slice5()
	data := []string{"red", "", "black", "", "", "pink", "blue", "orange", "gray", "purple", "yellow"}
	afterData := noEmpty(data)
	fmt.Printf("afterData:%v\n", afterData)

	s1 := afterData[6:] // [purple, yellow] 
	s2 := afterData[0:5] // "red", black", pink", "blue", "orange"
	copy(s2, s1)
	fmt.Printf("s2=%v\n", s2) // [purple yellow pink blue orange]
}