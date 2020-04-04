package main

import (
	"errors"
	"fmt"
	"math"
)

/**
	一. 使用 New 函数创建自定义错误
		创建自定义错误最简单的方法是使用 errors 包中的 New 函数。
	二. 使用 Errorf 给错误添加更多信息
	三. 使用结构体类型和字段提供错误的更多信息
	四. 使用结构体类型和方法，提供更多错误信息
 */

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func circleArea3(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

type areaError struct {
	err    string
	radius float64
}

type areaError2 struct {
	err    string //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError2{err, length, width}
	}
	return length * width, nil
}


func main() {
	//radius := -20.0
	//area, err := circleArea2(radius)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("Area of circle %0.2f", area)


	//radius := -20.0
	//area, err := circleArea3(radius)
	//if err != nil {
	//	if err, ok := err.(*areaError); ok {
	//		fmt.Printf("Radius %0.2f is less than zero", err.radius)
	//		//return
	//	}
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("Area of rectangle1 %0.2f", area)


	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}