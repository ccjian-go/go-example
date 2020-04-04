package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
	一. 创建一个有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
		同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
		通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
			ch := make(chan type, capacity)
		要让一个信道有缓冲，上面语法中的 capacity 应该大于 0。无缓冲信道的容量默认为 0
	二.	write 协程中向 ch 的写入发生了阻塞，直到 ch 有值被读取到。
		而 Go 主协程休眠了两秒后，才开始读取该信道，因此在休眠期间程序不会打印任何结果。
		主协程结束休眠后，在第 19 行使用 for range 循环，开始读取信道 ch，
		打印出了读取到的值后又休眠两秒，这个循环一直到 ch 关闭才结束。
	三. 我们向容量为 2 的缓冲信道写入 3 个字符串。当在程序控制到达第 3 次写入时（第 11 行），
		由于它超出了信道的容量，因此这次写入发生了阻塞。
		现在想要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。
	四. 长度 vs 容量
		缓冲信道的容量是指信道可以存储的值的数量。
		我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
		缓冲信道的长度是指信道中当前排队的元素个数。
		但在本例中，并没有并发协程来读取这个信道，因此这里会发生死锁（deadlock）。
		程序会在运行时触发 panic
	五. WaitGroup 用于实现工作池
		WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞
	六. WaitGroup 是一个结构体类型，我们创建了 WaitGroup 类型的变量，其初始值为零值。
		WaitGroup 使用计数器来工作。当我们调用 WaitGroup 的 Add 并传递一个 int 时，
		WaitGroup 的计数器会加上 Add 的传参。
		要减少计数器，可以调用 WaitGroup 的 Done() 方法。
		Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。
	七. 缓冲信道的重要应用之一就是实现工作池。
		工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配
	八. 工作池的核心功能
		创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
		将作业添加到该输入型缓冲信道中。
		作业完成后，再将结果写入一个输出型缓冲信道。
		从输出型缓冲信道读取并打印结果。
	九. 传递 wg 的地址是很重要的。如果没有传递 wg 的地址，那么每个 Go 协程将会
		得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道
*/
func write (ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
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

func worker23(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{
			job,
			digits23(job.randomno),
		}
		results <- output
	}
	wg.Done()
}

func createWorkerPool23(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker23(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate23(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)
	}

	//ch3 := make(chan string, 2)
	//ch3 <- "naveen"
	//ch3 <- "paul"
	//ch3 <- "steve"
	//fmt.Println(<-ch3)
	//fmt.Println(<-ch3)

	ch4 :=make(chan string, 3)
	ch4 <- "naveen"
	ch4 <- "paul"
	fmt.Println("capacity is", cap(ch4))
	fmt.Println("length is", len(ch4))
	fmt.Println("read value", <-ch4)
	fmt.Println("new length is", len(ch4))

	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")


	startTime := time.Now()
	noOfJobs := 100
	go allocate23(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool23(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
