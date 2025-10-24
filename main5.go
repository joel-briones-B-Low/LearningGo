package main

import "fmt"

// ! * -> pointer
// ! & -> address

func plus5(number *int) {
	*number++
}
func main5() {
	valor := 10
	println("before de plus:", valor)
	plus5(&valor)
	println("after de plus:", valor)

	pointer := new(int)
	fmt.Println("before de plus:", &pointer)
	*pointer = 20
	fmt.Println("before de plus:", *pointer)
}
