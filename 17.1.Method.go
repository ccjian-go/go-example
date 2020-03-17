package main

import (
	"fmt"
	"math"
)
/**
	一. 方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
		接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
	二. Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	三. 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
		假设我们有一个 Square 和 Circle 结构体。可以在 Square 和 Circle 上分别定义一个 Area 方法。
	四. changeName 方法有一个值接收器 (e Employee)，而 changeAge 方法有一个指针接收器 (e *Employee)。
		在 changeName 方法中对 Employee 结构体的字段 name 所做的改变对调用者是不可见的，
		因此程序在调用 e.changeName("Michael Andrew") 这个方法的前后打印出相同的名字。
		由于 changeAge 方法是使用指针 (e *Employee) 接收器的，
		所以在调用 (&e).changeAge(51) 方法对 age 字段做出的改变对调用者将是可见的
	五. 指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时
		当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在
		方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
		在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
	六. 属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段
	七. 当一个函数有一个值参数，它只能接受一个值参数。
		当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	八. 也可以在非结构体类型上定义方法
		为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。
		到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 main 包中，因此它们是可以运行的
		我们尝试把一个 add 方法添加到内置的类型 int。这是不允许的，因为 add 方法的定义和 int 类型的定义不在同一个包中
	九. 我们为 int 创建了一个类型别名 myInt
		我们定义了一个以 myInt 为接收器的的方法 add
 */

type Employee17 struct {
	name     string
	salary   int
	currency string
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e Employee17) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
	displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e Employee17) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}


type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}



type Employee171 struct {
	name string
	age  int
}

/*
使用值接收器的方法。
*/
func (e Employee171) changeName(newName string) {
	e.name = newName
}
/*
使用指针接收器的方法。
*/
func (e *Employee171) changeAge(newAge int) {
	e.age = newAge
}

// ===============================================================
type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("\nFull address: %s, %s \n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}


type myInt int
func (a myInt) add(b myInt) myInt {
	return a + b
}



func main() {
	emp1 := Employee17 {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法
	displaySalary(emp1)


	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())


	e := Employee171{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf( "Employee name before change: %s", e.name)
	e.changeName("Michael Andrew2")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
	e.changeAge(52)
	fmt.Printf("\nEmployee age after change: %d", e.age)


	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address {
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress() //访问 address 结构体的 fullAddress 方法


	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}


