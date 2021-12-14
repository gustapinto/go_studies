package main;

import (
	"fmt"
)

func main() {
	// Go has full C-like pointer support
	value := 7

	pointer := &value  // Point to value

	fmt.Println(pointer)  // See the memory block of the pointer
	fmt.Println(*pointer)  // See the value of the pointer

	*pointer = 14

	fmt.Println(value)  // See the new value of value
}
