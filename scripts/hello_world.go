package main

// Import packages, every go program and file is a set of packages
import (
	"fmt"
	"math/rand"
	"time"
)

// Declaring a function
func get_random_number() int {
		// Go has a deterministic nature, so even the random numbers are not
	// random until a diferent seed is set, like C and C++ this is
	// normaly set using a datetime object converted to a integer
	datetime := time.Now()  // Get the current datetime
	unix_time := datetime.Unix()  // Get a Int64 represantation of time

	rand.Seed(unix_time)

	return rand.Intn(1000)
}

// Declaring the entrypoint function
func main() {
	// On go you can declare variables usign a var keyword or a short
	// atribution operator, :=
	message := "Hello World !"

	fmt.Println(message)

	// And reassign a value of the same type usign the = operator
	message = "Random number"

	fmt.Println(message, get_random_number())
}
