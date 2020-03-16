package main

import (
	"fmt"
)

/**
	知识点
		一. 如何声明和使用可变函数
			1. 使用 变量...类型 声明
			2. 使用 变量... 传入切片
			3. 使用 变量1,变量2 可逐个个传入

		二. 如何可变函数的特点
			我们使用了语法糖 ... 并且将切片作为可变参数传入 change 函数。
			正如前面我们所讨论的，如果使用了 ... ，
			welcome 切片本身会作为参数直接传入，不需要再创建一个新的切片。

 */
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	find2(89, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func find2(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
