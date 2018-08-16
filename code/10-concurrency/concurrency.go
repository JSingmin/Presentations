// Code heavily based off golang's sample documentation: https://tour.golang.org

package main

import (
	"fmt"
	"time"
)

func basicExample() {
	printStuff := func(description string) {
		fmt.Printf("%v start\n", description)
		time.Sleep(2500 * time.Millisecond)
		fmt.Printf("%v end\n", description)
	}

	printStuff("main 1")

	for i := 1; i <= 5; i++ {
		go printStuff(fmt.Sprintf("goroutine %v", i))
	}

	printStuff("main 2")
	// time.Sleep(2000 * time.Millisecond)
}

func channelsExample() {
	triangularNumber := func(n int, c chan int) {
		fmt.Printf("%v start\n", n)
		time.Sleep(2000 * time.Millisecond)
		c <- (n * (n + 1)) / 2
		fmt.Printf("%v end\n", n)
	}

	var channel chan int = make(chan int)
	var n int = 5

	for i := 1; i <= n; i++ {
		go triangularNumber(i, channel)
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("result: %v\n", <- channel)
	}

	for i := 1; i <= n; i++ {
		go triangularNumber(i, channel)
		fmt.Printf("result: %v\n", <- channel)
	}
}

func main() {
	basicExample()
	// channelsExample()
}
