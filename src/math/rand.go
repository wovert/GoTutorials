package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 种子数: 种子数rand.Seed(1) 每次生成随机数相同
	rand.Seed(time.Now().UnixNano())
}

func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
