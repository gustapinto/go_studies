package main

import (
	"fmt"
)

// Go has support for full featured structs that ressamble the C/C++ type
// but with a class like functionality support
type Person struct {
	name string
	age uint
	secret string
}

// Remember that in go structs are OOP like? So here is how to declare
// a callable method within a struct
func (person *Person)ChangeAge(newAge uint) {
	// This function also uses a pointer to modify structs fields
	person.age = newAge  // Reassign a struct field value
}

func (person Person)Describe() {
	fmt.Println("Name: ", person.name, "Age: ", person.age)
}

func main() {
	// Create a struct instance without specifying the fields names
	mike := Person{"Mike", 21, "Dont know his real name"}

	// Create a struct instance with named fields
	bob := Person{
		age: 52,
		name: "Bob",
		secret: "Belongs to a cleaner dimension",
	}

	mike.ChangeAge(21)
	mike.Describe()

	// Creating a pointer to a struct, this is mostly used to pass
	// bigger structs by reference in function calls
	pointer := &bob

	(*pointer).age = 21 // Access the pointer filed the fancy way
	pointer.age = 20 // Access the pointer field the simple way

	bob.Describe()

	// Creating a anonymous struct without specifyng an named type
	// and with context only lifetime, it is mostrly used a temporary
	// data placeholder or tranformator
	desk := struct {
		material string
		foots int
	}{
		material: "iron",
		foots: 5,
	}

	fmt.Println(desk.material, desk.foots)
}
