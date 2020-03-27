package main

import (
	"fmt"
	"time"
)

func main() {
	countDown()
}

func countDown(){
	for i :=  3; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Go!")
			break
		}
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}