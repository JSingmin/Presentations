// Code heavily based off golang's sample documentation:
// https://tour.golang.org
// https://golang.org/doc/codewalk/functions/

package main

import (
	"fmt"
)

type calculation func(int, int) int

func add(a int, b int) int {
	return a + b
}

func compute(a int, b int, fn calculation) int {
	return fn(a, b)
}

func interceptor(fn calculation) calculation {
	return func(a int, b int) int {
		return fn(a, b) + 1
	}
}

func main() {
	a := 2
	b := 3
	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Printf("%v + %v = %v\n", a, b, compute(a, b, add))
	fmt.Printf("%v * %v = %v\n", a, b, compute(a, b, multiply))
	fmt.Printf("INTERCEPTED: %v * %v = %v\n", a, b, compute(a, b, interceptor(multiply)))
}
