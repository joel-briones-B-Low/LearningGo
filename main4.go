package main

import (
	"errors"
	"fmt"
)

func plus4(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

// funciont return more of one value
func slice4(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by 0")
	}
	cociente := a / b
	return cociente, nil
}

func printValues4(names ...string) {
	length := len(names)
	for i, value := range names {
		if length-1 == i {
			fmt.Println(value)
		} else {
			fmt.Print(value, ", ")
		}
	}
}

func count4() func() int {
	c := 0

	return func() int {
		c++
		return c
	}
}

type Rectangle4 struct {
	Width  float64
	Height float64
}

func (r Rectangle4) Area4() float64 {
	return r.Width * r.Height
}

func main4() {
	//valuePlus := plus(3, 4)
	//defer fmt.Println(valuePlus)
	//defer fmt.Println("Defer en main")
	//cociente, err := slice(10, 0)
	//if err != nil {
	//	println("Error:", err.Error())
	//} else {
	//	println("Cociente:", cociente)
	//}

	printValues4("Alice", "Bob", "Charlie")
	counter := count4()
	fmt.Println(counter())
	fmt.Println(counter())

	rect := Rectangle4{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", rect.Area4())

}
