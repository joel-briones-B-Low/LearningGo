package main

import "fmt"

type Shape6 interface {
	Area() float64
}

type Circle6 struct {
	Radius float64
}

func (c Circle6) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea6(s Shape6) {
	fmt.Printf("Area is : %.2f\n", s.Area())
}

func main6() {

	c := Circle6{Radius: 5}
	printArea6(c)

}
