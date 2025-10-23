package main

import (
	"fmt"
	"strings"
)

func main2() {
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
	onlyHigh := strings.ToUpper(mensage)
	fmt.Println("Lowercase Message:", onlyLow)
	fmt.Println("Uppercase Message:", onlyHigh)

	isAdult := age >= 18 // boolean
	fmt.Println("Is Adult:", isAdult)

	staticArray := [5]int{1, 2, 3} // static array
	fmt.Println("Length of Static Array:", len(staticArray))
	fmt.Println("Static Array:", staticArray)
	fmt.Println("Capacity of Static Array:", cap(staticArray))

	dynamicSlice := []int{}
	fmt.Println("Dynamic Slice before:", dynamicSlice)
	dynamicSlice = append(dynamicSlice, 10, 20, 30) // dynamic slice
	fmt.Println("Dynamic Slice:", dynamicSlice)
	fmt.Println("Length of Dynamic Slice:", len(dynamicSlice))
	fmt.Println("Capacity of Dynamic Slice:", cap(dynamicSlice))

	dic := map[string]int{"one": 1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Dictionary:", dic)

	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 28}
	fmt.Println("Person Struct:", p)

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
