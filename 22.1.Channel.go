package main

import (
	"fmt"
	"time"
)

/**
	一. 所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
		chan T 表示 T 类型的信道。
		信道的零值为 nil。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 make 来定义信道。
	二. 定义了一个 int 类型的信道 a
	三. 信道的发送和接收
		箭头对于 a 来说是向外指的，因此我们读取了信道 a  的值，并把该值存储到变量 data
			data := <- a // 读取信道 a
		箭头指向了 a ，因此我们在把数据写入信道 a
			a <- data // 写入信道 a
		信道旁的箭头方向指定了是发送数据还是接收数据。
	四. 发送与接收默认是阻塞的
		当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
		与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	五. 通过信道 done 接收数据。这一行代码发生了阻塞，除非有协程向 done 写入数据，否则程序不会跳到下一行代码。
		于是，这就不需要用以前的 time.Sleep 来阻止 Go 主协程退出了
	六. 每个函数都有传递信道的参数，以便写入数据。Go 主协程会在第 33 行等待两个信道传来的数据。
		一旦从两个信道接收完数据，数据就会存储在变量 squares 和 cubes 里，然后计算并打印出最后结果
	七. 使用信道需要考虑的一个重点是死锁。当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。
		如果没有的话，程序就会在运行时触发 panic，形成死锁。
		同理，当有 Go 协程等着从一个信道接收数据时，我们期望其他的 Go 协程会向该信道写入数据，
		要不然程序就会触发 panic。
	八. 单向通道
		我们目前讨论的信道都是双向信道，即通过信道既能发送数据，又能接收数据。
		其实也可以创建单向信道，这种信道只能发送或者接收数据。
	九. 这就需要用到信道转换（Channel Conversion）了。
		把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。
	十. 关闭信道
		数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
		当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
		v, ok := <- ch
			如果成功接收信道所发送的数据，那么 ok 等于 true。
			而如果 ok 等于 false，说明我们试图读取一个关闭的通道。
			从关闭的信道读取到的值会是该信道类型的零值。
			例如，当信道是一个 int 类型的信道时，那么从关闭的信道读取的值将会是 0。
	十一. for range 循环用于在一个信道关闭之前，从信道接收数据。
 */

func hello22(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}


func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func digits22(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares22(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits22(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes22(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits22(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}


func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello22(done)
	<-done // 等待读取
	fmt.Println("Main received data")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)

	//sendch := make(chan<- int)
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received ",v)
	}

	number22 := 589
	sqrch22 := make(chan int)
	cubech22 := make(chan int)
	go calcSquares22(number22, sqrch22)
	go calcCubes22(number22, cubech22)
	squares22, cubes22 := <-sqrch22, <-cubech22
	fmt.Println("Final output", squares22+cubes22)
}
