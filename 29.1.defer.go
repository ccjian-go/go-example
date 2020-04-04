package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	一. defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。
		这个定义可能看起来很复杂，我们通过一个示例就很容易明白了。
	二. defer 不仅限于函数的调用，调用方法也是合法的。
	三. 并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值。
	四. 当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，
		随后按照后进先出（Last In First Out, LIFO）的顺序执行
	五. 如果你仔细观察，会发现 wg.Done() 只在 area 函数返回的时候才会调用。
		wg.Done() 应该在 area 将要返回之前调用，并且与代码流的路径（Path）无关，
		因此我们可以只调用一次 defer，来有效地替换掉 wg.Done() 的多次调用。
	六. 使用 defer 还有一个好处。假设我们使用 if 条件语句，又给 area 方法添加了一条返回路径（Return Path）。
		如果没有使用 defer 来调用 wg.Done()，我们就得很小心了，确保在这条新添的返回路径里调用了 wg.Done()。
		由于现在我们延迟调用了 wg.Done()，因此无需再为这条新的返回路径添加 wg.Done()
 */

func finished(){
	time.Sleep(time.Second * 1)
	t := time.Now()   //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	fmt.Println("finished at " + t.Format("2006-01-02 15-04-05"))
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	t := time.Now()   //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	fmt.Println("Started at " + t.Format("2006-01-02 15-04-05"))
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}}
	fmt.Println("Largest number in", nums, "is", max)
}

type person29 struct {
	firstName string
	lastName string
}

func (p person29) fullName() {
	fmt.Printf(" - %s %s",p.firstName,p.lastName)
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func (r rect) area2(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	//nums := []int{78, 109, 2, 563, 300}
	//largest(nums)
	//
	//
	//p := person29 {
	//	firstName: "John",
	//	lastName: "Smith",
	//}
	//defer p.fullName()
	//fmt.Printf("Welcome ")
	//
	//
	//name := "Naveen"
	//fmt.Printf("Orignal String: %s\n", string(name))
	//fmt.Printf("Reversed String: ")
	//for _, v := range []rune(name) {
	//	defer fmt.Printf("%c", v)
	//}


	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		//go v.area(&wg)
		go v.area2(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}