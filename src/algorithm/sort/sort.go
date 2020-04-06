package main

import "fmt"

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
	for i:=1; i<len(a);i++ {
		for j:=1; j<len(a)-i; j++ {
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

	data := [...]int{8,3,4,2,1,9}
	isort(data[:])
	fmt.Println(data)
}