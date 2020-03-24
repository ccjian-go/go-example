package main

import (
	"fmt"
	"path/filepath"
)

/**
	一. 错误一直是很常见的。错误用内建的 error 类型来表示。
		就像其他的内建类型（如 int、float64 等），错误值可以存储在变量里、作为函数的返回值等等。
	二. 我们解析了这条错误信息，虽然获取了发生错误的文件路径，但是这种方法很不优雅。
		随着语言版本的更新，这条错误的描述随时都有可能变化，使我们程序出错。
	三.
		1. 断言底层结构体类型，使用结构体字段获取更多信息
		2. 断言底层结构体类型，调用方法获取更多信息
 */

type error interface {
	Error() string
}

func main() {
/*
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
*/

/*
	f2, err2 := os.Open("/test.txt")
	if err2, ok := err2.(*os.PathError); ok {
		fmt.Println("File at path", err2.Path, "failed to open")
		return
	}
	fmt.Println(f2.Name(), "opened successfully")
 */

/*
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
 */

	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)

}