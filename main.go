package main

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
}
