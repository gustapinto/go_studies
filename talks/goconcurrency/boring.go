package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; i < 10; i++ {
		// Pass the message to the channel, awaiting for the channel to receive
		// the value with a blocking behaviour
		c <- fmt.Sprintf("%s %d", msg, i)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// Generator is a pattern where functions return channels
func boringGenerator(msg string) <-chan string {
	c := make(chan string)

	// Opens a anon goroutine
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	// Creates a channel to comunicate data between goroutines
	c := make(chan string)

	// Executes boring as a goroutine
	go boring("boring", c)

	for i := 0; i < 10; i++ {
		// Uses the value of the channel to print the messages, awaiting the value
		fmt.Println(<-c)
	}

	c2 := boringGenerator("generator 1")
	c3 := boringGenerator("generator 2")

	// Because of the blocking nature of channel operators two generators
	// using the samechannel to print wiil print in sequence
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
		fmt.Println(<-c3)
	}
}
