// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
)

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Printf("index = %v; value = %v\n", i, v)
	}
}

func filterArray(arr []int, filter func(int) bool) []int {
	temp := make([]int, 0)

	for _, v := range arr {
		if filter(v) {
			temp = append(temp, v)
		}
	}

	return temp
}

func main() {
	var smallArray [2]int
	smallArray[0] = 1
	smallArray[1] = smallArray[0] + 1
	fmt.Printf("smallArray: %v; length = %v\n", smallArray, len(smallArray))

	largeArray := []int{1, 2, 3, 4, 5, 6}
	printArray(largeArray)

	printArray(filterArray(largeArray, func(n int) bool {
		return n >= 5
	}))
}
