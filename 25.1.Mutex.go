package main

import (
	"fmt"
	"sync"
)

/**
	一. 当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码称为临界区。
	二. 根据上下文切换的不同情形，x 的最终值是 1 或者 2。这种不太理想的情况称为竞态条件（Race Condition）
	三. 如果在任意时刻只允许一个 Go 协程访问临界区，那么就可以避免竞态条件。而使用 Mutex 可以达到这个目的。
	四. Mutex 用于提供一种加锁机制（Locking Mechanism），
		可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
	五. Mutex 可以在 sync 包内找到。Mutex 定义了两个方法：Lock 和 Unlock。
		所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。
	六. 第 7 行的 increment 函数把 x 的值加 1，并调用 WaitGroup 的 Done()，通知该函数已结束。
		在上述程序的第 15 行，我们生成了 1000 个 increment 协程。
		每个 Go 协程并发地运行，由于第 8 行试图增加 x 的值，因此多个并发的协程试图访问 x 的值，
		这时就会发生竞态条件
		字面意思 竞争状态的条件
	七. 使用 Mutex
		在前面的程序里，我们创建了 1000 个 Go 协程。如果每个协程对 x 加 1，最终 x 期望的值应该是 1000。
		在本节，我们会在程序里使用 Mutex，修复竞态条件的问题
	八. 使用信道处理竞态条件
		我们还能用信道来处理竞态条件
	九. 使用 容量为 1 的缓冲信道
		我们创建了容量为 1 的缓冲信道，并在第 18 行将它传入 increment 协程。
		该缓冲信道用于保证只有一个协程访问增加 x 的临界区。
		具体的实现方法是在 x 增加之前（第 8 行），传入 true 给缓冲信道。
		由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞，直到 x 增加后，
		信道的值才会被读取（第 10 行）。实际上这就保证了只允许一个协程访问临界区。
	十. 当 Go 协程需要与其他协程通信时，可以使用信道。而当只允许一个协程访问临界区时，可以使用 Mutex。
		就我们上面解决的问题而言，我更倾向于使用 Mutex，因为该问题并不需要协程间的通信。
		所以 Mutex 是很自然的选择
 */

var x int = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

var x2 int =0

func increment2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x2 = x2 + 1
	m.Unlock()
	wg.Done()
}

var x3 int  = 0
func increment3(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x3 = x3 + 1
	<- ch
	wg.Done()
}

func main(){
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)


	var w2 sync.WaitGroup
	var m2 sync.Mutex
	for i := 0; i < 1000; i++ {
		w2.Add(1)
		go increment2(&w2,&m2)
	}
	w2.Wait()
	fmt.Println("final value of x", x2)


	var w3 sync.WaitGroup
	ch3 := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w3.Add(1)
		go increment3(&w3, ch3)
	}
	w3.Wait()
	fmt.Println("final value of x", x3)
}