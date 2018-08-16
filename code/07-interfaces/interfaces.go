// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

type PrintableName interface {
	PrintName()
}

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintName() {
	fmt.Printf("%v %v\n", p.FirstName, p.LastName)
}

type Pet struct {
	Name string
}

func main() {
	var person PrintableName = Person{
		FirstName: "FN",
		LastName:  "LN",
	}
	person.PrintName()

	var pet PrintableName = Pet{
		Name: "Bob",
	}

	var pet2 interface{} = Pet{
		Name: "Bob2",
	}

	pet2 = 42
	pet2 = "OHAI"
	pet2 = false
}
