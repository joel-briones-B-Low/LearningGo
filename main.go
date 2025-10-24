package main

import (
	"errors"
	"fmt"
)

func plus(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

// funciont return more of one value
func slice(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by 0")
	}
	cociente := a / b
	return cociente, nil
}

func printValues(names ...string) {
	length := len(names)
	for i, value := range names {
		if length-1 == i {
			fmt.Println(value)
		} else {
			fmt.Print(value, ", ")
		}
	}
}

func count() func() int {
	c := 0

	return func() int {
		c++
		return c
	}
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	//valuePlus := plus(3, 4)
	//defer fmt.Println(valuePlus)
	//defer fmt.Println("Defer en main")
	//cociente, err := slice(10, 0)
	//if err != nil {
	//	println("Error:", err.Error())
	//} else {
	//	println("Cociente:", cociente)
	//}

	printValues("Alice", "Bob", "Charlie")
	counter := count()
	fmt.Println(counter())
	fmt.Println(counter())

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", rect.Area())

}
