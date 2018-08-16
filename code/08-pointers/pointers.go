// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint
}

func printPersonByValue(p Person, callingMethod string) {
	fmt.Printf("%v -> printPersonByValue (%v: %v)\n", callingMethod, p.Name, p.Age)
}

func printPersonByPointer(p *Person, callingMethod string) {
	fmt.Printf("%v -> printPersonByPointer (%v: %v)\n", callingMethod, p.Name, p.Age)
}

func (p Person) ItsMyBirthdayByValue() {
	p.Age += 1
	printPersonByValue(p, "ItsMyBirthdayByValue")
	printPersonByPointer(&p, "ItsMyBirthdayByValue")
}

func (p *Person) ItsMyBirthdayByPointer() {
	p.Age += 1
	printPersonByValue(*p, "ItsMyBirthdayByValue")
	printPersonByPointer(p, "ItsMyBirthdayByValue")
}

func valueAndPointerManipulationExample() {
	var ageValue uint = 30
	var agePointer *uint = &ageValue

	fmt.Printf("(ageValue: %v; agePointer: %v)\n", ageValue, agePointer)
	fmt.Printf("(&ageValue: %v; *agePointer: %v)\n", &ageValue, *agePointer)

	ageValue += 1
	fmt.Printf("(ageValue: %v; agePointer: %v)\n", ageValue, agePointer)
	fmt.Printf("(&ageValue: %v; *agePointer: %v)\n", &ageValue, *agePointer)

	*agePointer += 1
	fmt.Printf("(ageValue: %v; agePointer: %v)\n", ageValue, agePointer)
	fmt.Printf("(&ageValue: %v; *agePointer: %v)\n", &ageValue, *agePointer)
}

func valueFunctionsAndMethodsExample() {
	var personValue Person = Person{
		Name: "personValue",
		Age:  767,
	}

	fmt.Printf("Original: %v %v\n", personValue.Name, personValue.Age)
	personValue.ItsMyBirthdayByValue()
	fmt.Printf("After ItsMyBirthdayByValue: %v %v\n", personValue.Name, personValue.Age)
	personValue.ItsMyBirthdayByPointer()
	fmt.Printf("After ItsMyBirthdayByPointer: %v %v\n", personValue.Name, personValue.Age)
}

func pointerFunctionsAndMethodsExample() {
	var personPointer *Person = &Person{
		Name: "personPointer",
		Age:  498,
	}

	fmt.Printf("Original: %v %v\n", personPointer.Name, personPointer.Age)
	personPointer.ItsMyBirthdayByValue()
	fmt.Printf("After ItsMyBirthdayByValue: %v %v\n", personPointer.Name, personPointer.Age)
	personPointer.ItsMyBirthdayByPointer()
	fmt.Printf("After ItsMyBirthdayByPointer: %v %v\n", personPointer.Name, personPointer.Age)
}

func main() {
	valueAndPointerManipulationExample()
	// valueFunctionsAndMethodsExample()
	// pointerFunctionsAndMethodsExample()
}
