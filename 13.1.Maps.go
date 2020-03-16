package main

import (
	"fmt"
)

/**
	一. make方法初始化
	二. 直接声明时初始化
	三. 获取 map 元素的语法是 map[key]
	四. 如果获取一个不存在的元素，map 会返回该元素类型的零值
	五. 如果我们想知道 map 中到底是不是存在这个 key，value, ok := map[key]
		如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
	六. 当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同
	七. 任意变量类型都可使用 printf + %T + 变量 输出变量的类型
	八. 删除 map 中 key 的语法是 delete(map, key)。这个函数没有返回值。
	九. 获取 map 的长度使用 len 函数
	十. 和 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。
		因此，改变其中一个变量，就会影响到另一变量。
 */
func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	personSalary2 := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary2["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary2)

	personSalary3 := map[string]int{}
	fmt.Println("personSalary map contents:", personSalary3)

	personSalary4 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary4[employee])
	fmt.Println("Salary of joe is", personSalary4["joe"])

	newEmp := "joe"
	value, ok := personSalary4[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}

	fmt.Println("All items of a map")
	for key, value := range personSalary4 {
		fmt.Printf("personSalary4[%s] = %d\n", key, value)
	}

	fmt.Printf("type of personSalary4 is %T", personSalary4)

	fmt.Println("map before deletion", personSalary4)
	delete(personSalary4, "steve")
	fmt.Println("map after deletion", personSalary4)

	fmt.Println("length is", len(personSalary4))

	newPersonSalary := personSalary4
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary4)
}