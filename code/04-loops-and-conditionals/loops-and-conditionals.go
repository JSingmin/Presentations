// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

func printEvenForLoop(start int, end int) {
	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			fmt.Printf("Even numbers for-loop: %v\n", i)
		}
	}
}

func printOddWhileLoop(start int, end int) {
	i := start
	for i <= end {
		if remainder := i % 2; remainder == 1 {
			fmt.Printf("Odd numbers for-loop: %v (remainder %v)\n", i, remainder)
		}

		i += 1
	}
}

func main() {
	defer fmt.Println("Done")

	printEvenForLoop(1, 5)
	printOddWhileLoop(1, 5)
}
