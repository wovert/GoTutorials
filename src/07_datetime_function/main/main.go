package main

import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

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

	// 时间常量
	i := 0
	for {
		i++
		fmt.Println(i)
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i == 10 {
			break
		}
	}

	// Unix 和 UnixNano
	fmt.Printf("Unix时间戳=%v UnixNano时间戳=%v\n", now.Unix(), now.UnixNano())

	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	extime := end - start
	fmt.Printf("执行test耗费时间为%v秒\n", extime)
}



