package main

import "fmt"

// 声明 add_func函数变量
type add_func func(int, int) int

func add(a,b int) int {
	return a + b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func modify(a int) {
	a = 100
}

func sub(a,b int) (c int) {
	c = a - b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a+b)/2
	return
}

func a() {
	i := 0
	defer fmt.Println("defer:", i) // 最后执行 defer
	i++
	fmt.Printf("i=%d", i)
	return
}

func f() {
	for i:=0; i<5; i++ {
		defer fmt.Printf("%d", i)
	}
}

func myfunc(args ...int) {
	fmt.Println("len(args) =", len(args))
	for i, data := range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}

func main() {
	// c := add
	// fmt.Println(c)

	// sum := c(10, 20)
	// fmt.Println(sum)

	// res := operator(c, 100, 200)
	// fmt.Println(res)

	// res2 := operator(sub, 100 ,20)
	// fmt.Println(res2)

	// res3, res4 := calc(100, 200)
	// fmt.Println(res3) // 300
	// fmt.Println(res4) // 150

	// a()

	// j := new(int)
	// *j = 100
	// fmt.Println(j)

	var a []int
	a = append(a, 10, 20, 392)
	a = append(a, a...)
	fmt.Println(a) 


	myfunc(10, 20, 30)
	
}