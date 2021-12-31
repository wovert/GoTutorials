package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func slice0() {
	intArr := [...]int{1, 2, 3, 4, 5, 6}
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
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//s := arr[1:3:5]
	s := arr[1:3] // 6-1=5 cap=元素数量-开始解决索引值

	fmt.Println("s=", s)           // [2 3]
	fmt.Println("len(s)=", len(s)) // 2
	fmt.Println("cap(s)", cap(s))  // 5
}

func slice2() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[1:5] // cap默认是8

	fmt.Println("s=", s)           // [2 3 4 5]
	fmt.Println("len(s)=", len(s)) // 4
	fmt.Println("cap(s)", cap(s))  // 8-1=7
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
	var slice []float64 = []float64{1.0, 2.0, 3.0}
	fmt.Printf("slice=%v\n", slice)
	slice[2] = 10.0
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("len(slice)=%v\n", len(slice))
	fmt.Printf("cap(slice)=%v\n", cap(slice))
}

func slice5() {
	s1 := []int{1, 2, 4, 5}
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

func fb(n int) []uint64 {
	fbSlice := make([]uint64, n)
	fbSlice[0] = 1
	fbSlice[1] = 1

	for i := 2; i < n; i++ {
		fbSlice[i] = fbSlice[i-1] + fbSlice[i-2]
	}
	return fbSlice
}

func main() {
	var names []string

	// 类型 []string
	fmt.Printf("%T\n", names)

	// 零值 []string(nil)
	fmt.Printf("%#v\n", names)

	// 初始化
	// 字面量

	// 第一种切片赋值
	names = []string{} // 空切片 已经初始化但是元素数量为 0
	fmt.Printf("%#v\n", names)

	names = []string{"vue", "react", "angular"}
	fmt.Printf("%#v\n", names) // []string{"vue", "react", "angular"}

	// 第二种切片赋值
	names = []string{1: "vue", 3: "php"}
	fmt.Printf("%#v\n", names) // []string{"", "vue", "", "php"}

	// 第三种 make函数
	/**
	 	make() 2个参数
		make() 3个参数
		struct {
			*array 内存中数组的指针
			length 存储元素的数量
			cap 容量（<=1024元素 length*2, > 1024元素 length*0.25）
		}
	*/
	names = make([]string, 5)                 // 字符串切片类型，5个元素 并0值初始化
	fmt.Printf("%#v\n", names)                // []string{"", "", "", "", ""}
	fmt.Printf("len(names)=%d\n", len(names)) // len(names)=5
	fmt.Printf("cap(names)=%d\n", cap(names)) // cap(names)=5
	names[0] = "vue"
	names[2] = "NodeJS"
	// names[6] = "php" // panic: runtime error: index out of range [6] with length 5
	fmt.Printf("%#v\n", names) // []string{"", "", "", "", ""}
	fmt.Println("---------")

	names = make([]string, 0, 4)
	fmt.Printf("%#v\n", names)                // []string{}
	fmt.Printf("len(names)=%d\n", len(names)) // len(names)=0
	fmt.Printf("cap(names)=%d\n", cap(names)) // cap(names)=4
	fmt.Println("---------")

	// names[0] = "vue"  panic: runtime error: index out of range [0] with length 0
	names = append(names, "vue")
	fmt.Printf("%#v\n", names)                // []string{"vue"}
	fmt.Printf("len(names)=%d\n", len(names)) // len(names)=0
	fmt.Printf("cap(names)=%d\n", cap(names)) // cap(names)=4
	fmt.Println("---------")

	names = append(names, "js")
	names = append(names, "html")
	names = append(names, "css")
	names = append(names, "php")
	fmt.Printf("%#v\n", names)                // []string{"vue", "js", "html", "css", "php"}
	fmt.Printf("len(names)=%d\n", len(names)) // len(names)=5
	fmt.Printf("cap(names)=%d\n", cap(names)) // cap(names)=8
	fmt.Println("---------")

	// 遍历
	for k, v := range names {
		fmt.Printf("names[%d]=%#v\n", k, v)
	}
	fmt.Println("---------")

	// copy 切片之间赋值
	aSlice := []string{"a", "b", "c"}
	bSlice := []string{"x", "y", "z"}
	// copy(目的，源) 长度相等
	copy(aSlice, bSlice)        // aSlice复制到bSlice
	fmt.Printf("%#v\n", aSlice) // []string{"x", "y", "z"}
	fmt.Println("---------")

	cSlice := []string{"a", "b", "c"}
	dSlice := []string{"x", "y", "z", "o"}
	// copy(目的，源) 源切片元素数量大于目的切片元素数量
	copy(cSlice, dSlice)
	// 如果源切片的元素数量大于目的切片元素数量，赋值之后以目的元素数量为结果
	fmt.Printf("%#v\n", cSlice) // []string{"x", "y", "z"}
	fmt.Println("---------")

	eSlice := []string{"a", "b", "c"}
	fSlice := []string{"x", "y"}
	// copy(目的，源) 源切片元素数量小于目的切片元素数量
	copy(eSlice, fSlice)
	// 如果源切片的元素数量小于目的切片元素数量，赋值之后多余源切片元素的保留响应元素索引位置
	fmt.Printf("%#v\n", eSlice) // []string{"x", "y", "c"}
	fmt.Println("---------")

	// 切片操作 names[start:end]
	nums := []int{0, 1, 2, 3, 4, 5}
	slice0 := nums[1:3]                         // start <= end
	fmt.Printf("%#v\n", slice0)                 // []int{1, 2}
	fmt.Printf("len(slice0)=%d\n", len(slice0)) // len(slice0)=2
	fmt.Printf("cap(slice0)=%d\n", cap(slice0)) // cap(slice0)=5
	fmt.Println("---------")

	slice1 := nums[3:3]
	fmt.Printf("%#v\n", slice1)                 // []int{}
	fmt.Printf("len(slice1)=%d\n", len(slice1)) // len(slice1)=0
	fmt.Printf("cap(slice1)=%d\n", cap(slice1)) // cap(slice1)=3
	fmt.Println("---------")

	// start<= end <= cap
	slice2 := nums[2:6]
	fmt.Printf("%#v\n", slice2)                 // []int{2,3,4,5}
	fmt.Printf("len(slice2)=%d\n", len(slice2)) // len(slice2)=4
	fmt.Printf("cap(slice2)=%d\n", cap(slice2)) // cap(slice2)=4
	fmt.Println("---------")

	nums = make([]int, 0, 10)
	slice3 := nums[2:3]
	fmt.Printf("slice3=%#v\n", slice3)          // []int{0, 100}
	fmt.Printf("len(slice3)=%d\n", len(slice3)) // len(slice3)=2
	// 现cap=源cap-源start
	fmt.Printf("cap(slice3)=%d\n", cap(slice3)) // cap(slice3)=8
	fmt.Println("---------")

	nums = []int{0, 1, 2, 3, 4, 5}
	slice4 := nums[2:3]
	fmt.Printf("nums=%#v\n", nums) // nums=[]int{0, 1, 2, 3, 4, 5}
	fmt.Println("---------")

	slice4 = append(slice4, 100)
	fmt.Printf("slice4=%#v\n", slice4)          // []int{2, 100}
	fmt.Printf("len(slice4)=%d\n", len(slice4)) // len(slice4)=2
	// 现cap=源cap-源start
	fmt.Printf("cap(slice4)=%d\n", cap(slice4)) // cap(slice4)=4
	fmt.Println("---------")

	fmt.Printf("nums=%#v\n", nums)          // int{0, 1, 2, 100, 4, 5}
	fmt.Printf("len(nums)=%d\n", len(nums)) // len(nums)=6
	fmt.Printf("cap(nums)=%d\n", cap(nums)) // cap(nums)=6
	fmt.Println("---------")

	nums = []int{0, 1, 2, 3, 4, 5}
	// [start:end:max] start<=end<=max<=cap, new_cap=max-start
	slice5 := nums[2:3:6]
	slice5 = append(slice5, 100)
	fmt.Printf("nums=%#v\n", nums)          // []int{2, 100}
	fmt.Printf("len(nums)=%d\n", len(nums)) // len(slice5)=2
	// 现cap=源cap-源start
	fmt.Printf("cap(nums)=%d\n", cap(nums)) // cap(slice5)=4
	fmt.Println("---------")

	fmt.Printf("slice5=%#v\n", slice5)        // int{0, 1, 2, 100, 4, 5}
	fmt.Printf("len(nums)=%d\n", len(slice5)) // len(slice5)=6
	fmt.Printf("cap(nums)=%d\n", cap(slice5)) // cap(slice5)=6
	fmt.Println("---------")

	// 移除第一个元素
	slice6 := nums[1:]
	fmt.Printf("nums=%#v\n", nums)
	fmt.Printf("slice6=%#v\n", slice6)

	// 移除最后一个元素
	slice7 := nums[:len(nums)-1]
	fmt.Printf("slice7=%#v\n", slice7)

	// 移除中间的元素
	// []int{0, 1, 2, 100, 4, 5}
	// []int{0, 1, 4, 5}
	target := nums[2:] // 目标
	copy(target, nums[:2])
	fmt.Printf("源=%#v\n", nums[:2]) // []int{0, 1}
	fmt.Printf("目标=%#v\n", target)  // []int{0, 1, 4, 5}
	fmt.Printf("nums=%#v\n", nums)  // []int{0, 1, 0, 1, 4, 5}

	// 队列：先进先出

	return

	// slice0()
	//slice1()
	// slice2()
	// slice3()
	// slice4()
	//slice5()
	//return

	a1 := [...]int{1, 3, 5, 7, 9}
	ss1 := a1[0:4]
	fmt.Println(ss1)
	fmt.Printf("len(ss1)=%d, cap(ss1)=%d", len(ss1), cap(ss1)) // 4, 5

	data := []string{"red", "", "black", "", "", "pink", "blue", "orange", "gray", "purple", "yellow"}
	afterData := noEmpty2(data)
	fmt.Printf("afterData:%v\n", afterData)

	// s1 := afterData[6:] // [purple, yellow]
	// s2 := afterData[0:5] // "red", black", pink", "blue", "orange"
	// copy(s2, s1)
	// fmt.Printf("s2=%v\n", s2) // [purple yellow pink blue orange]

	// strTest()

	// fbSlice := fb(10)
	// fmt.Println("fbSlice=", fbSlice)

	// append

	x1 := [...]int{1, 3, 5} // array
	x2 := x1[:]             // slice
	fmt.Println(x1, len(x1), cap(x2))
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &x2[0])
	x2 = append(x2[:1], x2[2:]...) // 修改了底层数组
	fmt.Printf("%p\n", &x2[0])
	fmt.Println(x2, len(x2), cap(x2))
	x2[0] = 100     // 修改了底层数组
	fmt.Println(x1) // [100 5 5]

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
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
		s1 = append(s1, 2*i+1)
	}
	fmt.Println(s1)

	// create slice method
	s2 := []int{2, 4, 6, 7, 8, 10, 29}
	printSlice(s2)

	// 16:len, 20:cap
	s3 := make([]int, 16)

	// 32: cap
	s4 := make([]int, 10, 32)

	printSlice(s3)
	printSlice(s4)

	fmt.Println("=========== Copy slice ===========")
	copy(s3, s2)   // s3复制到s2
	printSlice(s3) // [2 4 6 7 0 0 0 0 0....]

	fmt.Println("=========== Delete elemes from slice ===========")
	s3 = append(s3[:3], s3[4:]...) // delete 7, ...展开符
	printSlice(s3)                 // [2 4 6 0 0 0 0 0....]

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
