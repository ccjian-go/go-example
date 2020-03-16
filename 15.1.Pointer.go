package main

import (    "fmt")

/**
	一. 指针是一种存储变量内存地址（Memory Address）的变量。
	二. 指针变量的类型为 *T，该指针指向一个 T 类型的变量。
	三. & 操作符用于获取变量的地址
	四. 指针的零值是 nil
	五. 指针的解引用 可以 获取 指针所指向的变量的 值。将 a 解引用的语法是 *a。
	六. 向函数传递指针参数，在函数内使用解引用，可以修改外部值
	七. 不要向函数传递数组的指针，而应该使用切片
	八. a[x] 是 (*a)[x] 的简写形式
	九. 向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。
		我们最好使用切片来处理。
	十. Go 不支持指针运算

 */
func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	a1 := 25
	var b1 *int
	if b1 == nil {
		fmt.Println("b is", b1)
		b1 = &a1
		fmt.Println("b after initialization is", b1)
	}

	b2 := 255
	a2 := &b2
	fmt.Println("address of b is", a2)
	fmt.Println("value of b is", *a2)

	a3 := 58
	fmt.Println("value of a before function call is",a3)
	b3 := &a3
	unChange(a3)
	fmt.Println("value of a after function call is", a3)
	change2(b3)
	fmt.Println("value of a after function call is", a3)

	a4 := [3]int{89, 90, 91}
	modify(&a4)
	fmt.Println(a4)

	a5 := [3]int{89, 90, 91}
	modify2(&a5)
	fmt.Println(a5)

	a6 := [3]int{89, 90, 91}
	modify3(a6[:])
	fmt.Println(a6)

	//b7 := [...]int{109, 110, 111}
	//p := &b7
	//p++
}

func change2(val *int) {
	*val = 55
}

func unChange(val int) {
	val = 55
}

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(arr *[3]int) {
	arr[0] = 90
}

func modify3(sls []int) {
	sls[0] = 90
}