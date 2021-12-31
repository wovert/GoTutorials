package main

import (
	"fmt"
	"sort"
)

//func BubbleSort(arr *[]int) {
//	fmt.Println("排序前arr=", (*arr))
//	fmt.Println("len(arr)=", len(*arr))
//	// 第一轮
//	temp := 0
//	for i := 0; i < len(*arr) - 1; i++ {
//		cnt := len(*arr) - 1 - i
//		for j := 0; j < cnt; j++ {
//			if (*arr)[j] > (*arr)[j+1] {
//				temp = (*arr)[j]
//				(*arr)[j] = (*arr)[j+1]
//				(*arr)[j+1] = temp
//			}
//		}
//		fmt.Printf("第%v轮排序后arr=%v\n", (i + 1), (*arr))
//	}
//}
//

func isort(a []int) {
	fmt.Println("排序前arr=", a)
	fmt.Println("len(arr)=", len(a))
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
		fmt.Printf("第%v轮排序后arr=%v\n", i, a)
	}
}
func main() {
	//arr := []int{74, 68, 50, 57, 12}
	//BubbleSort(&arr)

	data := [...]int{8, 3, 4, 2, 1, 9}
	isort(data[:])
	fmt.Println(data)

	nums := []int{8, 3, 4, 2, 1, 9}
	fmt.Printf("排序前 nums=%#v\n", nums)
	sort.Ints(nums)
	fmt.Printf("排序后 nums=%#v\n", nums)

	names := []string{"abc", "xyz", "mn", "k", "z"}
	fmt.Printf("排序前 names=%#v\n", names)
	sort.Strings(names)
	fmt.Printf("排序后 names=%#v\n", names)

	// 倒序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Printf("排序后 nums=%#v\n", nums)
}
