package main

import (
	"fmt"
	"time"
)

func main() {
	// Defer, this postpone the call just before the end of a function,
	// they can be stacked in a stack like data structure
	defer fmt.Println("I must be the first")

	// A classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	// Go uses For to represent a while loop
	j := 1

	for j < 10 {
		fmt.Println("j: ", j)

		j += j
	}

	// A if else block
	falsy := false

	if falsy {
		fmt.Println("Truly")
	} else {
		fmt.Println("Falsy")
	}

	// Like a for loop a If declaretion can had a shorthand declaration
	if n := 10; n <= j {
		fmt.Println("He is")
	}

	// Flip a Switch, because they work like C switchs but without the
	// need for a break statement
	a := "b"

	switch a {
		case "c":
			fmt.Println("C")
		default:
			fmt.Println("A")
	}

	// In Go switchs can have no condition, always validating to true
	// This strategy is use to create a cleaner sintax for long if else
	// blocks, like on a factory pattern
	hour := time.Now().Hour()

	switch {
		case hour < 12:
			fmt.Println("Good morning")
		case hour < 18:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good night")
	}
}
