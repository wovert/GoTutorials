package main

import (
	"fmt"
)

func BubbleSort2(a *[]int) {
	fmt.Println("排序前", *a)
	cnt := len(*a)
	temp := 0
	for i := 0; i < cnt - 1; i++ {
		fmt.Printf("第%v轮排序%v\n", (i+1), *a)
		for j := 0; j < cnt - 1 - i; j++ {
			if (*a)[j] > (*a)[j+1] {
				temp = (*a)[j]
				(*a)[j] = (*a)[j-1]
				(*a)[j-1] = temp
			}
		}
	}
	fmt.Println("排序后", *a)
}
func main() {
	arr := []int{20,12,9,72,32,52}
	BubbleSort2(&arr)
}