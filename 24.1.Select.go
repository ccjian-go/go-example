package main

import (
	"fmt"
	"time"
)

/**
	一. select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
		如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。
		该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。
	二. 假设我们有一个关键性应用，需要尽快地把输出返回给用户。
		这个应用的数据库复制并且存储在世界各地的服务器上。
		假设函数 server1 和 server2 与这样不同区域的两台服务器进行通信。
		每台服务器的负载和网络时延决定了它的响应时间。
		我们向两台服务器发送请求，并使用 select 语句等待相应的信道发出响应。
		select 会选择首先响应的服务器，而忽略其它的响应。
		使用这种方法，我们可以向多个服务器发送请求，并给用户返回最快的响应了。:）
	三. 在没有 case 准备就绪时，可以执行 select 语句中的默认情况（Default Case）。
		这通常用于防止 select 语句一直阻塞。
	四. 如果存在默认情况，就不会发生死锁，因为在没有其他 case 准备就绪时，会执行默认情况。
	五. 当 select 由多个 case 准备就绪时，将会随机地选取其中之一去执行。
	六. select {}
		除非有 case 执行，select 语句就会一直阻塞着。
		在这里，select 语句没有任何 case，因此它会一直阻塞，导致死锁
	七.
 */

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
	}


	ch2 := make(chan string)
	select {
	case <-ch2:
	default:
		fmt.Println("default case executed")
	}

	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
			case v := <-ch:
				fmt.Println("received value: ", v)
				return
			default:
				fmt.Println("no value received")
		}
	}


}