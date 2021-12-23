package main

import (
	"fmt"
)

/* Go implements concurrency in the form of goroutines, they are
 * tiny concurrent subprocesses that are much smaller and lightweight
 * when compared to a classic thread, because of this is commom to use
 * a LOT of goroutines in a tipical Go program
 */

func foo(done chan bool) {
	fmt.Println("Foo")

	// Goroutines can be declared inside another routine, and also as
	// a anonimous function
	a := make(chan bool)

	go func() {
		fmt.Println("Bar")

		close(a)
	}()

	<-a

	done <- true
}

/* It's important to know that main() also executes in a goroutine, the
 * "main routine" and it will be exit inependent of other goroutines, so
 * we need to use a atrategy to mantain a goroutine alive, tipical a chan
 */
func main() {
	/* Goroutines automatic call return when they are instantiated, so
	 * for the function to run we need to define a channel, and use this
	 * channel to mantain the goroutine active. Channels act like Pipes
	 * and are used to pass data between goroutines thus preventing race
	 * conditions
	 */
	done := make(chan bool)  // Defines a tipical "done" control channel

	go foo(done)  // Defines a goroutine that use "done" to be kept alive

	<-done
}
