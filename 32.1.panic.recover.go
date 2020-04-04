package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

/**
	一. 在有些情况，当程序发生异常时，无法继续运行。
		在这种情况下，我们会使用 panic 来终止程序。
		当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
		这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，
		然后程序会打印出 panic 信息，		接着打印出堆栈跟踪（Stack Trace），最后程序终止
	二. 当程序发生 panic 时，使用 recover 可以重新获得对该程序的控制。
	三. 你应该尽可能地使用错误，而不是使用 panic 和 recover。只有当程序不能继续运行的时候，
		才应该使用 panic 和 recover 机制。
	四. 发生了一个不能恢复的错误，此时程序不能继续运行。
		一个例子就是 web 服务器无法绑定所要求的端口。在这种情况下，就应该使用 panic，
		因为如果不能绑定端口，啥也做不了。
	五. 发生了一个不能恢复的错误，此时程序不能继续运行。
		一个例子就是 web 服务器无法绑定所要求的端口。
		在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。
	六. 发生了一个编程上的错误。假如我们有一个接收指针参数的方法，而其他人使用 nil 作为参数调用了它。
		在这种情况下，我们可以使用 panic，因为这是一个编程错误：用 nil 参数调用了一个只能接收合法指针的方法。
	七. 内建函数 panic 的签名如下所示：
		func panic(interface{})
	八. 当程序终止时，会打印传入 panic 的参数。

	九. 发生 panic 时的 defer
		当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
		这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，
		接着打印出堆栈跟踪，最后程序终止。
	十. recover 是一个内建函数，用于重新获得 panic 协程的控制。
		func recover() interface{}
	十一. 只有在延迟函数的内部，调用 recover 才有用。
		在延迟函数内调用 recover，可以取到 panic 的错误信息，并且停止 panic 续发事件（Panicking Sequence），
		程序运行恢复正常。如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
	十二. 只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic。
	十三. 运行时 panic
		运行时错误（如数组越界）也会导致 panic。
		这等价于调用了内置函数 panic，其参数由接口类型 runtime.Error 给出。
		runtime.Error 接口的定义如下：
		type Error interface {
			error
			// RuntimeError is a no-op function but
			// serves to distinguish types that are run time
			// errors from ordinary errors: a type is a
			// run time error if it has a RuntimeError method.
			RuntimeError()
		}
	十四. 当我们恢复 panic 时，我们就释放了它的堆栈跟踪。
		实际上，在上述程序里，恢复 panic 之后，我们就失去了堆栈跟踪。
		有办法可以打印出堆栈跟踪，就是使用 Debug 包中的 PrintStack 函数。
*/

func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	//defer fmt.Println("deferred call in fullName")

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}


func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}


func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func aa() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	//defer fmt.Println("deferred call in main")
	//firstName := "Elon"
	//fullName(&firstName, nil)
	//fmt.Println("returned normally from main")

	//a()
	//fmt.Println("normally returned from main")

	//defer r()
	aa()
	fmt.Println("normally returned from main")
}