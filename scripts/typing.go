package main

import (
	"fmt"
)

// Beside its basic type Go also offer support to higher level types, like
// Structs
type person struct {
	name string
	age  uint
}

// Interfaces
type downloader interface {
	download() bool
	prepare() string
}

// Functions, because they can be first class types in Go
type foo func() bool

func main() {
	// Go is a static typed language, so it has its types, Go basic
	// types are:

	// Strings, to represent a character and character arrays
	message := "Foo Bar"
	char := "A"

	// Booleans
	truly := true
	falsy := false

	/* A f***** ton of integers that I cant even think of variable names
	 * so it will be a comment block
	 *
	 * int  int8  int16  int32  int64
	 * uint uint8 uint16 uint32 uint64 uintptr
	 */
	int_number := 999

	/* Some floats
	 *
	 * float32 float64
	 */
	float_number := 999.999

	/* And complex numbers, with real and imaginary parts
	 *
	 * complex64 complex128
	 */
	complex_number := complex(10, 11)

	// On Go you can also type hint a variable type usign the var keyword
	var boolean bool = true
	var integer uint = 99999
	var text string = "FOOBAR"

	fmt.Println(message, char)
	fmt.Println(truly, falsy, int_number)
	fmt.Println(float_number, real(complex_number), imag(complex_number))
	fmt.Println(boolean, integer, text)
}
