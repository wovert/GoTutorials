package main

import (
	"fmt"
	"sort"
)

func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	mid := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		fmt.Printf("找不到\n")
		return
	}
	if (*arr)[mid] > findVal {
		BinarySearch(arr, leftIndex, mid-1, findVal)
	} else if (*arr)[mid] < findVal {
		BinarySearch(arr, mid+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下表为%v\n", mid)
	}
}

func main() {
	/**
	二分查找分析：arr = [1, 8, 10, 90, 100, 2932]
	问题：查找一个从小到大的有序数组中的数字
	分析：
		1. 找到数组中的中间下标：mid = (leftIndex + rightIndex) / 2，然后让中间下标的值和findVal 进行比较
		2.1 如果 arr[mid] > findVal, leftIndex = mid - 1
		2.2 如果 arr[mid] < findVal, mid+1 = rightIndex
		2.3 如果 arr[mid] == findVal
		3. 2.1-2.2-2.3 逻辑递归执行
		4. 退出递归条件：leftIndex > rightIndex
	*/
	arr := [6]int{1, 8, 10, 90, 100, 2932}

	BinarySearch(&arr, 0, len(arr)-1, -2932)

	nums := []int{1, 8, 10, 90, 100, 2932}
	fmt.Printf("search 90 is %t\n", nums[sort.SearchInts(nums, 90)] == 90)
}
