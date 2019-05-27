package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	fmt.Printf("Year=%v\n", now.Year())
	fmt.Printf("Month=%v\n", int(now.Month()))
	fmt.Printf("Day=%v\n", now.Day())
	fmt.Printf("Hour=%v\n", now.Hour())
	fmt.Printf("Minute=%v\n", now.Minute())
	fmt.Printf("Seconds%v\n", now.Second())

	// 格式化日期时间
	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d \n", now.Year(),
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr:", dateStr)

	fmt.Println(now.Format("2006/01/02 15:04:05"))
}



