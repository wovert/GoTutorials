package main

import (
	"fmt"
	"math"
)

// 水仙花数：一个三位数，其各位数字立方和等于该数本身
func isNumber(n int) bool {
	var i, j, k int
	i = n % 10
	j = (n / 10) % 10
	k = (n / 100) % 10
	sum := i*i*i + j*j*j + k*k*k

	return sum == n
}

func isPrime(n int) bool {
	// fmt.Printf("sqrt:%d\n", int(math.Sqrt(float64(n))))
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 阶乘
func sum(n int) uint64 {

	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s := s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		sum += s
	}
	return sum
}

func main() {
	var n int
	var m int

	fmt.Scanf("%d,%d", &n, &m)

	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}

	for i := n; i < m; i++ {
		if isNumber(i) == true {
			fmt.Println("水仙花数：", i)
		}
	}

	s := sum(n)
	fmt.Println(s)

	letters := "爱上基地决赛"
	for _, v := range letters {
		fmt.Printf("%q\n", v)
	}

}
