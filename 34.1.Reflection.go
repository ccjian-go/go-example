package main

import (
	"fmt"
	"reflect"
)

/**
	一. 反射就是程序能够在运行时检查变量和值，求出它们的类型。
	二. 如果程序中每个变量都是我们自己定义的，那么在编译时就可以知道变量类型了，
		为什么我们还需要在运行时检查变量，求出它的类型呢？
	三. reflect 包
		reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层具体类型和具体值。
		这正是我们所需要的。createQuery函数接收 interface{} 参数，根据它的具体类型和具体值，
		创建 SQL 查询。这正是 reflect 包能够帮助我们的地方。
	四. reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值
	五. reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别
		返回 reflect.Type 和 reflect.Value
		这两种类型是我们创建查询生成器的基础
	六. reflect.Kind
		reflect 包中还有一个重要的类型：Kind
		Type 表示 interface{} 的实际类型（在这里是 main.Order)，
		而 Kind 表示该类型的特定类别（在这里是 struct）
	七. NumField() 和 Field() 方法
		NumField() 方法返回结构体中字段的数量，而 Field(i int) 方法返回字段 i 的 reflect.Value
	八. Int() 和 String() 方法
		Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string
	九. 使用了 Name() 方法，从该结构体的 reflect.Type 获取了结构体的名字
	十. 使用反射编写清晰和可维护的代码是十分困难的。你应该尽可能避免使用它，只在必须用到它时，才使用反射。
 */

type order struct {
    ordId      int
    customerId int
}

type employee struct {
    name string
    id int
    address string
    salary int
    country string
}

func createQuery(q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Type's name is", t.Name())
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
	return ""
}


func createQuery2(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
				case reflect.Int:
					if i == 0 {
						query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
					} else {
						query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
					}
				case reflect.String:
					if i == 0 {
						query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
					} else {
						query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
					}
				default:
					fmt.Println("Unsupported type")
					return
			}

			query = fmt.Sprintf("%s", query)
		}
		query += ")"
		fmt.Println(query)
		//fmt.Println("unsupported type")
	}
}


func main() {
	//o := order{
	//	ordId:      456,
	//	customerId: 56,
	//}
	//createQuery(o)
	//
	//
	//a := 56
	//x := reflect.ValueOf(a).Int()
	//fmt.Printf("type:%T value:%v\n", x, x)
	//b := "Naveen"
	//y := reflect.ValueOf(b).String()
	//fmt.Printf("type:%T value:%v\n", y, y)


	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery2(e)
	//i := 90
	//createQuery2(i)
}