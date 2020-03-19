package main
import (
	"fmt"
)

/**
	一. 在 Go 语言中，接口就是方法签名（Method Signature）的集合。
		当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。
		接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
	二. 在 Go 中，并不需要一个类使用 implement 关键字，来显式地声明该类实现了接口
		一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口
	三. 接口的实际用途
		totalExpense 可以扩展新的员工类型，而不需要修改任何代码。
		假如公司增加了一种新的员工类型 Freelancer，它有着不同的薪资结构。
		Freelancer只需传递到 totalExpense 的切片参数中，无需 totalExpense 方法本身进行修改。
		只要 Freelancer 也实现了 SalaryCalculator 接口，totalExpense 就能够实现其功能。
	四. 可以把接口看作内部的一个元组 (type, value)。
		type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
	五. 没有包含方法的接口称为空接口。空接口表示为 interface{}。
		由于空接口没有方法，因此所有类型都实现了空接口。
		describe(i interface{}) 函数接收空接口作为参数，因此，可以给这个函数传递任何类型
	六. 类型断言用于提取接口的底层值（Underlying Value）。
		在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
		v, ok := i.(T)
		如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
		如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。
	七. 类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。
		唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。
		类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，
		而对于类型选择，类型 T 由关键字 type 代替。
		还可以将一个类型和接口相比较。
		如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较

 */
// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId  int
	basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
	total expense is calculated by iterating though the SalaryCalculator slice and summing
	the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}


type Test interface {
	Tester()
}
type MyFloat float64
func (m MyFloat) Tester() {
	fmt.Println(m)
}
func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	//s := i.(int) //get the underlying int value from i
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch i.(type) {
		case string:
				fmt.Printf("I am a string and my value is %s\n", i.(string))
			case int:
				fmt.Printf("I am an int and my value is %d\n", i.(int))
			default:
				fmt.Printf("Unknown type\n")
	}
}

type Describer4 interface {
	Describe4()
}

type Person4 struct {
	name string
	age  int
}
func (p Person4) Describe4() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType4(i interface{}) {
	switch v := i.(type) {
		case Describer4:
			v.Describe4()
		default:
			fmt.Printf("unknown type\n")
	}
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()

	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)

	var i2 interface{} = 56
	assert(i2)
	var s2 interface{} = "Steven Paul"
	assert(s2)

	findType("Naveen")
	findType(77)
	findType(89.98)

	findType4("Naveen")
	p := Person4{
		name: "Naveen R",
		age:  25,
	}
	findType4(p)
}