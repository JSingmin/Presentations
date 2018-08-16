// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"errors"
	"fmt"
)

func divide(n int, d int) (float32, error) {
	fmt.Printf("DIVIDE START\n")
	defer fmt.Printf("DIVIDE FINALLY\n")

	if (d == 0) {
		return 0, errors.New("denominator zero")
	}

	return float32(n / d), nil
}

func main() {
	if result, err := divide(10, 0); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("SUCCESS: %v\n", result)
	}
}
