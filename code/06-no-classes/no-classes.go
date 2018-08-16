// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintName() {
	fmt.Printf("%v %v\n", p.FirstName, p.LastName)
}

func main() {
	person := Person{
		FirstName: "FN",
		LastName:  "LN",
	}

	person.PrintName()
}
