package main


import (
	"fmt"
)

func fact(n int) int {
	if (n == 0) {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var num int = 5 // 5*4*3*2*1
	result := fact(num)
	fmt.Printf("!%d=%d\n", num, result)
}