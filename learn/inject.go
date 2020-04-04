package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	三. 依赖注入的好处
		测试代码。
			如果你不能很轻松地测试函数，这通常是因为有依赖硬链接到了函数或全局状态。
			例如，如果某个服务层使用了全局的数据库连接池，这通常难以测试，并且运行速度会很慢。
			DI 提倡你注入一个数据库依赖（通过接口），然后就可以在测试中控制你的模拟数据了。
		关注点分离，解耦了数据到达的地方和如何产生数据。
			如果你感觉一个方法 / 函数负责太多功能了（生成数据并且写入一个数据库？
			处理 HTTP 请求并且处理业务级别的逻辑），那么你可能就需要 DI 这个工具了。
		在不同环境下重用代码。
			我们的代码所处的第一个「新」环境就是在内部进行测试。
			但是随后，如果其他人想要用你的代码尝试点新东西，他们只要注入他们自己的依赖就可以了。
	四. 什么是模拟？我听说 DI 要用到模拟，它可讨厌了
		模拟（mocking）会在后面详细讨论（它并不坏）。
		你会使用模拟来代替真实事物，用一个模拟版本来注入，于是可以控制和检查你的测试。
		在我们的例子中，标准库已经有工具供我们使用了。
	五. GO 标准库
		通过熟悉 io.Writer 接口，我们可以用测试中的 bytes.Buffer 来作为 Writer，
		然后我们可以使用标准库中的其它的 Writer，在命令行应用或 web 服务器中使用这个函数。
		随着你越来越熟悉标准库，你就会越了解这些在代码中重用的通用接口，它们会使你的软件在许多场景都可以重用。
 */

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


