package main

import (
  "fmt"
//   "time"
  "runtime"
)

func test_routine() {
	for i :=1; i <= 10; i++{
		fmt.Printf("子协程执行%d\n", i)
		// time.Sleep(time.Second)
	}
}

func test() {
	defer fmt.Println("test function finish.")
	// return // 终止此函数
	runtime.Goexit() // 终止所在的协程
	fmt.Println("Excute test function.")
}

func main() {
	// 并行计算的CPU核数的最大值
	n := runtime.GOMAXPROCS(1)
	fmt.Printf("n = %d\n", n)

	// 指定CPU核数越少CPU交换次数越少
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}


	go test_routine()
	for i :=1; i <= 5; i++ {
		// 让出 CPU 时间片，先让别的协程执行，它执行完，再回来执行此协协程
		runtime.Gosched()

		fmt.Printf("主协程执行循环%d\n", i)
		
	}

	go func() {
		fmt.Println("匿名函数：子协程开始")
		
		test() // 函数里执行终止所在的协程

		// 因为总之所在协程，不会执行一下语句
		fmt.Println("匿名函数：子协程结束")
	}()

	// 主协程执行结束，其他所有协程都结束
	// time.Sleep(time.Second) // 1s
	// 死循环，不让主协程结束
	for {

	}
}