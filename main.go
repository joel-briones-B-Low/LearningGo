package main

import "fmt"

func main() {
	// conditionals if else ifelse
	edad := 20
	/*if edad < 18 {
		println("Eres menor de edad")
	} else if edad >= 18 && edad < 65 {
		println("Eres un adulto")
	} else {
		println("Eres un adulto mayor")
	*/

	if edad < 18 {
		println("Eres menor de edad")
		return
	}

	println("Eres un adulto")
	println("Fin del programa")

	// bucle for
	for i := 0; i < 5; i++ {
		println("Iteración número:", i)
		fmt.Printf("Iteración número: %d\n", i)
	}

}
