package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
	一. 无缓冲信道的发送和接收过程是阻塞的。
	二. 创建一个有缓冲（Buffer）的信道。
		只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
	三. 只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据
	四. 通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
		ch := make(chan type, capacity)
	五. 要让一个信道有缓冲，上面语法中的 capacity 应该大于 0。
		无缓冲信道的容量默认为 0，因此我们在创建无缓冲信道时，省略了容量参数
	六. 例3 - 由于它超出了信道的容量，因此这次写入发生了阻塞
		要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。
		但在本例中，并没有并发协程来读取这个信道，因此这里会发生死锁（deadlock）。
	七. 缓冲信道的容量是指信道可以存储的值的数量。
		我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
		缓冲信道的长度是指信道中当前排队的元素个数。
	八. 容量cap(ch)  长度len(ch)
	九. 而 WaitGroup 用于实现工作池，因此要理解工作池，我们首先需要学习 WaitGroup。
		WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
		假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。
		Go 主协程需要等待这 3 个协程执行结束后，才会终止。这就可以用 WaitGroup 来实现。
	十. WaitGroup 是一个结构体类型，我们创建了 WaitGroup 类型的变量，其初始值为零值。
		WaitGroup 使用计数器来工作。
	十一. 当我们调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。
		要减少计数器，可以调用 WaitGroup 的 Done() 方法。
		Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。
	十二. 传递 wg 的地址是很重要的。
		如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，
		因而当它们执行结束时，main 函数并不会知道。
	十三. 缓冲信道的重要应用之一就是实现工作池。
		一般而言，工作池就是一组等待任务分配的线程。
		一旦完成了所分配的任务，这些线程可继续等待任务的分配。
	十四. 工作池的核心功能如下：
		创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
		将作业添加到该输入型缓冲信道中。
		作业完成后，再将结果写入一个输出型缓冲信道。
		从输出型缓冲信道读取并打印结果。
	十五.
 */

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process23(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()

}

type Job struct {
	id       int
	randomno    int
}
type Result struct {
	job         Job
	sumofdigits     int
}
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

/**
	allocate 函数接收所需创建的作业数量作为输入参数，生成了最大值为 998 的伪随机数，
	并使用该随机数创建了 Job 结构体变量。
	这个函数把 for 循环的计数器 i 作为 id，最后把创建的结构体变量写入 jobs 信道。
	当写入所有的 job 时，它关闭了 jobs 信道。
 */
func allocate23(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

/**
	每个results遍历出来，便
 */
func result23(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

/**
	工作池
	每次创建一个工作 => 阻塞数+1
	在 go worker 协程工作完成之前
	wait一直阻塞等待完成
	每个worker 协程工作完成 => result channel便得到一个 result 结果

 */
func createWorkerPool23(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker23(&wg)
	}
	wg.Wait()
	close(results)
}

func worker23(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits23(job.randomno)}
		results <- output
	}
	wg.Done()
}

/**
digits 函数的任务实际上就是计算整数的每一位之和，最后返回该结果。
为了模拟出 digits 在计算过程中花费了一段时间，我们在函数内添加了两秒的休眠时间
*/
func digits23(number int) int{
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)

	return sum
}

func main() {
	//ch := make(chan string, 2)
	//ch <- "naveen"
	//ch <- "paul"
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)


	//ch := make(chan int, 2)
	//go write(ch)
	//time.Sleep((2 * time.Second))
	//for v := range ch {
	//	fmt.Println("read value", v,"from ch")
	//	time.Sleep(2 * time.Second)
	//}


	//ch := make(chan string, 2)
	//ch <- "naveen"
	//ch <- "paul"
	//ch <- "steve" // 已满被阻塞 deadlock 除非前面有 <-ch
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)


	//ch := make(chan string, 3)
	//ch <- "naveen"
	//ch <- "paul"
	//fmt.Println("capacity is", cap(ch))
	//fmt.Println("length is", len(ch))
	//fmt.Println("read value", <-ch)
	//fmt.Println("new length is", len(ch))


	//no := 3
	//var wg sync.WaitGroup
	//for i := 0; i < no; i++ {
	//	wg.Add(1)
	//	go process23(i, &wg)
	//}
	//wg.Wait()
	//fmt.Println("All go routines finished executing")


	startTime := time.Now()
	noOfJobs := 100
	go allocate23(noOfJobs)
	done := make(chan bool)
	go result23(done)
	noOfWorkers := 10
	createWorkerPool23(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}
