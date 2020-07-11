package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // 全局条件变量
const (
	COUNT = 3
)
// 生产者：发送数据端
func producer8(out chan<- int, index int) {
	for { // 循环目的是一个生产者不停的进行生产
		cond.L.Lock()			// 条件变量对应互斥加锁
		for len(out) == COUNT { // 产品区满等待消费者消费，缓冲区已经满了，为什么循环？
			cond.Wait()		// 阻塞(挂起)当前协程，等待条件变量满足，被消费者唤醒
			// 释放已掌握的互斥锁相当于 cond.L.Unlock, 两步为一个原子操作
			// 当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁，相当于 cond.L.Lock()
			// Wait()操作步骤
			// 1. 阻塞
			// 2. Unlock
			// 3. lock
		}
		num := rand.Intn(800) // 产生随机数
		out <- num // 写入到 channel 中 (生产)
		fmt.Printf("生产者%dth, 产生数据%3d, 公共区剩余%d个数据\n", index, num, len(out))
		cond.L.Unlock() // 生产结束，解锁互斥锁

		cond.Signal()		// 唤醒阻塞的消费者
		time.Sleep(time.Second) // 生产者休息一会，给其他协程执行机会
	}
}

// 消费者：接受数据端
func consumer8(in <- chan int, index int) {
	for {
		cond.L.Lock()	// 条件变量对应互斥加锁
		for len(in) == 0 {		// 产品区为空等待生产者生产
			cond.Wait()					// 挂起当前协程，等待条件变量满足，被生产者唤醒
		}
		num := <-in						// 将 channel 中的数据读走（消费）
		fmt.Printf("----消费者%dth, 消费数据%3d，公共区域剩余%d个数据\n", index, num, len(in))
		cond.L.Unlock()		// 消费结束，解锁互斥锁
		cond.Signal()			// 唤醒阻塞的生产者，单发通知，给一个等在（阻塞）在该条件变量上的 goroutine发送通知
		time.Sleep(time.Millisecond * 500)	// 消费完休息一会，给其他协程执行机会
	}
}

func main() {

	rand.Seed(time.Now().UnixNano()) 		// 设置随机种子
	quit := make(chan bool)							// 创建用于结束通信的 channel

	product := make(chan int, COUNT)		// 产品区（公共区）使用channel 模拟
	cond.L = new(sync.Mutex)						// 创建互斥锁，赋值在条件变量属性上


	// 5个生产者
	for i:=1; i<6; i++{
		go producer8(product, i)
	}


	// 3个消费者
	for i:=1; i<4; i++{
		go consumer8(product, i)
	}

	<-quit // 主协程，不结束
}
