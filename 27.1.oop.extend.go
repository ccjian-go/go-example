package main
/**
	一. 一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样。
		所以上面第 11 行的 p.author.fullName()  可以替换为 p.fullName() 。
	二. 嵌套的结构体切片 []post。错误的原因是结构体不能嵌套一个匿名切片。我们需要一个字段名
 */
import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title     string
	content   string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

func (p post) details2() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	//author1 := author{
	//	"Naveen",
	//	"Ramanathan",
	//	"Golang Enthusiast",
	//}
	//post1 := post{
	//	"Inheritance in Go",
	//	"Go supports composition instead of inheritance",
	//	author1,
	//}
	//post1.details2()

	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}

