package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Learning Go!")

	var age int = 25 // explicit type declaration
	fmt.Println("Age:", age)
	newAge := 30.5 // type inference
	fmt.Println("New Age:", newAge)

	mensage := "Welcome to Go programming!"
	newMensage := "Hello" + mensage
	fmt.Println(newMensage)

	total := age + int(newAge) // type conversion
	fmt.Println("Total Age:", total)

	onlyLow := strings.ToLower(mensage)
	fmt.Println("Lowercase Message:", onlyLow)

}

/*
data types in Go:
int int8 int16 int32 int64 // whole numbers | control the size
uint uint8 uint16 uint32 uint64 // unsigned whole numbers | only positive
float32 float64 // decimal numbers | precision control | depend on system
bool // true or false | flag or condition
string // sequence of characters | text
complex64 complex128  // complex numbers | real and imaginary parts
byte // alias for uint8 | represents a single byte of data
rune // alias for int32 | represents a unique character (Unicode code point)
*/
