package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/**
	一. 将字符串写入文件
		创建文件
		将字符串写入文件
	二. 将字节写入文件
		创建文件
		将字节写入文件
	三. 将字符串一行一行的写入文件
		创建文件
		循环一行行写入文件
	四. 追加到文件
		打开文件
		追加到文件末尾
	五. 并发写文件
		当多个 goroutines 同时（并发）写文件时，我们会遇到竞争条件(race condition)。
		因此，当发生同步写的时候需要一个 channel 作为一致写入的条件。
			创建一个 channel 用来读和写这个随机数。
			创建 100 个生产者 goroutine。每个 goroutine 将产生随机数并将随机数写入到 channel 里。
			创建一个消费者 goroutine 用来从 channel 读取随机数并将它写入文件。这样的话我们就只有一个 goroutinue 向文件中写数据，从而避免竞争条件。
			一旦完成则关闭文件。
		写一个例子，并发100个生产者协程去写，使用阻塞去等待完成，确认100写任务成功完成

 */

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main1() {
	f, err := os.Create("string.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := "Hello World2"
	l, err := f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}


	f2, err := os.Create("bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte {1, 2, 3, 4, 5}
	n2, err := f2.Write(d2)
	if err != nil {
		fmt.Println(err)
		f2.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")


	f3,err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f3.Close()
		return
	}
	d := []string{
		"Welcome to the world of Go1.",
		"Go is a compiled language.",
		"It is easy to learn Go.",
	}
	for _,v := range d{
		_,err = fmt.Fprintln(f3,v)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
	err = f3.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")


	f4, err := os.OpenFile("lines",os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		f3.Close()
		return
	}
	newLine := "File handling is easy."
	_,err = fmt.Fprintln(f4,newLine)
	if err !=nil {
		fmt.Println(err)
		f3.Close()
		return
	}
	err = f4.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

}

func main() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i<100; i++{
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data,done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	}else{
		fmt.Println("File writing failed")
	}

}




