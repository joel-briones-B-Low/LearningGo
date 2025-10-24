package main

import "fmt"

// ! * -> pointer
// ! & -> address

func plus(number *int) {
	*number++
}
func main() {
	valor := 10
	println("before de plus:", valor)
	plus(&valor)
	println("after de plus:", valor)

	pointer := new(int)
	fmt.Println("before de plus:", &pointer)
	fmt.Println("before de plus:", *pointer)
}
