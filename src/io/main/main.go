package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年两")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试")
	// fmt.Scanln(&isPass)

	fmt.Println("请输入姓名，年龄，薪水，是否通过考试")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	fmt.Printf("name is %v\n age is %v\n salary is %v\n pass %v", name, age, sal, isPass)
}