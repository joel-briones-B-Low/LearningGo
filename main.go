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

	// while
	n := 0
	for n < 5 {
		if n == 3 {
			println("Se encontró el valor 3, se salta la iteración")
			n++
			continue
		}
		if n == 4 {
			println("Se encontró el valor 4, se termina el bucle")
			break
		}

		println("While Iteración número:", n)
		n++
	}

	// infinite loop
	j := 0
	for {
		if j == 3 {
			println("Bucle infinito")
			break
		}
		if j%2 == 0 {
			println("Número par, se salta la iteración:", j)
			j++
			continue // skip even numbers and their actions
		}
		j++
	}

}
