package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)
	// p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	p.FirstName = "Eric yan"
	//p.SetFirstName("Eric")
	fmt.Println(p.FirstName) // Output: Eric

	fmt.Println(p.GetFirstName()) // Output: Eric
}