package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

/**
	一. 调用 fmt.Printf 会打印到标准输出，用测试框架来捕获它会非常困难。
		我们所需要做的就是注入（这只是一个等同于「传入」的好听的词）打印的依赖。
		我们的函数不需要关心在哪里打印，以及如何打印，所以我们应该接收一个接口，而非一个具体的类型。
	二. fmt.Printf 的源码，你可以发现一种引入（hook in）的方式
		func Printf(format string, a ...interface{}) (n int, err error) {
			return Fprintf(os.Stdout, format, a...)
		}
		在 Printf 内部，只是传入 os.Stdout，并调用了 Fprintf
		os.Stdout 究竟是什么？Fprintf 期望第一个参数传递过来什么？
		func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
			p := newPrinter()
			p.doPrintf(format, a)
			n, err = w.Write(p.buf)
			p.free()
			return
		}
		io.Writer 是：
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
 */

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer,"Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

//func Greet(writer *bytes.Buffer, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	//Greet(os.Stdout, "Elodie")
	Greet(os.Stdout, "Elodie")
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		fmt.Println(err)
	}
}


