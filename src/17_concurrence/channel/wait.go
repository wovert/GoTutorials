package main

import (
	"fmt"
	"sync"
)

// 比如我们有一个主任务在执行，执行到某一点时需要并行执行三个子任务，
//并且需要等到三个子任务都执行完后，再继续执行主任务。
//那我们就需要设置一个检查点，使主任务一直阻塞在这，等三个子任务执行完后再放行。
// WaitGroup 用来阻塞主协程，可以等待所有协程执行完
/**
Add 方法用于设置 WaitGroup 的计数值，可以理解为子任务的数量
Done 方法用于将 WaitGroup 的计数值减一，可以理解为完成一个子任务
Wait 方法用于阻塞调用者，直到 WaitGroup 的计数值为0，即所有子任务都完成
*/
// 总共三个方法Add(n)【n为总共要等待的协程数】，Done【在协程中调，相当于Add(-1)】Wait【等待阻塞结束】
func main() {
	// 读写锁
	var mutet sync.Mutex
	counter := 0
	// 适合一对多的 goroutine 的协同等待
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		// 等待 i 个 goroutine
		wg.Add(i)
		go func() {
			mutet.Lock()
			defer mutet.Unlock()
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
