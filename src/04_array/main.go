package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 值传递修改函数局部函数数组的元素值
func test1(arr [3]int) {
	arr[0] = 88
}

// 数组指针形式修改元素值
func test2(arr *[3]int) {
	(*arr)[0] = 88
}

func test3() {
	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100) // 0<=n<100
	}
	fmt.Println("交换前intArr=", intArr)

	temp := 0
	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后intArr=", intArr)
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}          // 初始化方式1
	arr3 := [...]int{2, 4, 6, 8, 10} // 初始化方式2
	var grid [4][5]bool
	strs := [3]string{1: "tom", 0: "jack", 2: "marry"} // 根据索引初始化

	fmt.Println(arr1, arr2, arr3) // [0 0 0 0 0] [1 3 5] [2 4 6 8 10]
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	for _, v := range strs {
		fmt.Println(v)
	}

	nums := [3]int{1, 2, 3}
	test1(nums)
	fmt.Println("nums=", nums)
	test2(&nums)
	fmt.Println("nums=", nums)

	test3()
}
