package main

import "fmt"

/*
	一. 结构体是用户定义的类型，表示若干个字段（Field）的集合。
		有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
	二. 通过把相同类型的字段声明在同一行，结构体可以变得更加紧凑。
	三. 声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 匿名结构体
	四. 当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值。
		 string类型变量的零值（"" ）。而int类型变量的零值（0）
		还可以为某些字段指定初始值，而忽略其他字段，被忽略的字段会赋值为零值。
		可以创建零值的 struct，以后再给各个字段赋值
	五. 点号操作符 . 用于访问结构体的字段
	六. 可以创建指向结构体的指针
	七. 我们在访问 firstName 字段时，可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName
	八. 我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段
		虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型
		比如在上面的 Person 结构体里，虽说字段是匿名的，但 Go 默认这些字段名是它们各自的类型。
		所以 Person 结构体有两个名为 string 和 int 的字段。
	九. 结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体。
	十. 如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。
		这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。
		Person 结构体有一个匿名字段 Address，而 Address 是一个结构体。
		现在结构体 Address 有 city 和 state 两个字段，
		访问这两个字段就像在 Person 里直接声明的一样，因此我们称之为提升字段。

 */
type Employee struct {
	firstName string
	lastName  string
	age       int
}

type Employee2 struct {
	firstName, lastName string
	age, salary         int
}

var employee3 struct {
	firstName, lastName string
	age int
}

type Person struct {
	string
	int
}

type Address2 struct {
	city, state string
}

type Person2 struct {
	name string
	age int
	address2 Address2
}

type Address3 struct {
	city, state string
}
type Person3 struct {
	name string
	age  int
	Address3
}


type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int //exported field
}


func main() {
	//creating structure using field names
	emp1 := Employee2{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee2{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	var emp4 Employee2 //zero valued structure
	fmt.Println("Employee 4", emp4)

	emp5 := Employee2{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	emp6 := Employee2{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
	fmt.Println("")

	var emp7 Employee2
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	emp8 := &Employee2{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	emp9 := &Employee2{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp9.firstName)
	fmt.Println("Age:", emp9.age)

	p := Person{"Naveen", 50}
	fmt.Println(p)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	var p2 Person2
	p2.name = "Naveen"
	p2.age = 50
	p2.address2 = Address2 {
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:",p2.age)
	fmt.Println("City:",p2.address2.city)
	fmt.Println("State:",p2.address2.state)

	var p3 Person3
	p3.name = "Naveen"
	p3.age = 50
	p3.Address3 = Address3{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city) //city is promoted field
	fmt.Println("State:", p3.state) //state is promoted field


	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}

