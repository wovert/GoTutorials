package main

import (
	"fmt"
	"sort"
)

func main() {
	stats := [][]int{{'A',3},{'B',2},{'C',1},{'D',2}}	

	// 使用出现次数进行排序
	// {B, 2}, {D, 2} => 稳定的（切片顺序）
	// {D, 2}, {B, 2} => 不稳定的

	// 不稳定的
	sort.Slice(stats, func(i, j int) bool { return stats[i][1] > stats[j][1]})
	fmt.Println(stats)

	// 稳定的
	sort.SliceStable(stats, func(i, j int) bool { return stats[i][1] > stats[j][1]})
	fmt.Println(stats)

	// 升序 <=
	// 降序 >=
	index := sort.Search(len(stats), func(i int) bool { return stats[i][1] == 1})
	fmt.Println(index) // 3
}