package main

import (
	"fmt"
)

func main() {
	// 1. 定义数组
	// 2. 控制体接受一个名词
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入查找的人名...")
	fmt.Scanln(&heroName)
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v, 下标%v\n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v \n", heroName)
		}
	}

	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("找到%v, 下标%v\n", heroName, index)
	}

}