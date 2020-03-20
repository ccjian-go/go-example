package main

import "fmt"

/**
	一. 使用值接受者声明的方法，既可以用值来调用，也能用指针调用。
		不管是一个值，还是一个可以解引用的指针，调用这样的方法都是合法的。
		对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
		但接口中存储的具体值（Concrete Value）并不能取到地址，
		因此在d2 = a并且d2.Describe()，对于编译器无法自动获取a 的地址，于是程序报错。
	二. 类型可以实现多个接口
		我们把e赋值给了SalaryCalculator类型的接口变量 ，同样把e赋值给LeaveCalculator类型的接口变量。
		由于e的类型Employee实现了SalaryCalculator和LeaveCalculator两个接口，因此这是合法的
	三. 接口嵌套
		Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口
		接口EmployeeOperations，它嵌套了两个接口：SalaryCalculator和LeaveCalculator
	四. 接口零值
		接口的零值是nil。
		对于值为nil的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为nil
 */

type Describer interface {
	Describe()
}
type Person191 struct {
	name string
	age  int
}

func (p Person191) Describe() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}


type SalaryCalculator19 interface {
	DisplaySalary()
}

type LeaveCalculator19 interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations19 interface {
	SalaryCalculator19
	LeaveCalculator19
}

type Employee19 struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e Employee19) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee19) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	var d1 Describer
	p1 := Person191{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person191{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe()


	e := Employee19 {
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 5000,
		pf: 200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator19 = e
	s.DisplaySalary()
	var l LeaveCalculator19 = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	var empOp EmployeeOperations19 = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

	var d19 Describer
	if d19 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d19, d19)
	}
}