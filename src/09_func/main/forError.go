package main


import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("for before", i)
		defer func(v int) {
			fmt.Println("defer", i) // i是最后的值
			fmt.Println("defer", v) // 2,1,0
		}(i)
		fmt.Println("for after", i)
	}

	fmt.Println("main")
}