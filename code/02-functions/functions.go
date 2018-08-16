// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func separateNumber(n float32) (integer int, remainder float32) {
	integer = int(n / 1)
	remainder = n - float32(integer)
	return

	// return int(n / 1), n - (float32(int(n / 1)))
}

func main() {
	fmt.Println(add(1, 1))

	integer, remainder := separateNumber(3.14)
	fmt.Println(integer)
	fmt.Println(remainder)
}
