package main

import (
	"fmt"
)

/* In Go the interfaces are used to enforce a struct to had a set
 * of methods. In Go interfaces are not implemented like on other OOP
 * languages, but inferred on a struct method creator
 */
type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  uint
}

type Dog struct {
	name string
}

// Make both Person and Dog compatible with the Describer interface by
// implementing his methods
func (person Person) Describe() {
	fmt.Println("Name: ", person.name, "Age: ", person.age)
}

func (dog Dog) Describe() {
	fmt.Println("Oof", dog.name)
}

// Doing this we can now create newly code that use the abstraction type
// instead of a concrete instance of the struct
func Handle(describer Describer) {
	describer.Describe()
}

func main() {
	mike := Person{"Mike", 21}
	doggo := Dog{"Doggo"}

	Handle(mike)
	Handle(doggo)
}
