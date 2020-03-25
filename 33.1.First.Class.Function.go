package main

import "fmt"

/**
	一. 什么是头等函数？
		支持头等函数（First Class Function）的编程语言，
		可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。
		Go 语言支持头等函数的机制。
	二. 由于没有名称，这类函数称为匿名函数（Anonymous Function）
	三. 要调用一个匿名函数，可以不用赋值给变量
	四. 用户自定义的函数类型
		type add func(a int, b int) int
	五. 我们定义了一个 add 类型的变量 a，并向它赋值了一个符合 add 类型签名的函数。
		我们调用了该函数，并将结果赋值给 s
	六. wiki 把高阶函数（Hiher-order Function）定义为：
		满足下列条件之一的函数：
				接收一个或多个函数作为参数
				返回值是一个函数
	七. 闭包（Closure）是匿名函数的一个特例。
		当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包
		每一个闭包都会绑定一个它自己的外围变量（Surrounding Variable）
	八. 头等函数的实际用途
 */

type add func(a int, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName string
	grade string
	country string
}
func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
	fmt.Println()
	func() {
		fmt.Println("hello world first class function2")
	}()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")


	var aa add = func(a int, b int) int {
		return a + b
	}
	s := aa(5, 6)
	fmt.Println("Sum", s)


	f := func(a, b int) int {
		return a + b
	}
	simple(f)


	s2 := simple2()
	fmt.Println(s2(60, 7))


	aaa := 5
	func() {
		fmt.Println("aaa =", aaa)
	}()


	a2 := appendStr()
	b2 := appendStr()
	fmt.Println(a2("World"))
	fmt.Println(b2("Everyone"))
	fmt.Println(a2("Gopher"))
	fmt.Println(b2("!"))


	s11 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s22 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	ss := []student{s11, s22}
	ff := filter(ss, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(ff)


	aaaa := []int{5, 6, 7, 8, 9}
	r := iMap(aaaa, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}