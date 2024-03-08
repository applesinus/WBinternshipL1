package main

import (
	"fmt"
)

// square function
func square3(x int, ch chan int) {
	// sending square of x to channel
	ch <- x * x
}

func Ex3() {
	fmt.Printf("\n==========  Exercise 3:  ==========\n\n")

	// creating a slice with given values
	slc := []int{2, 4, 6, 8, 10}

	// creating a buffered channel with the same length as the slice
	ch := make(chan int, len(slc))
	defer close(ch)

	// creating a loop which goes through the slice
	for _, val := range slc {
		// calling square function in goroutines
		go square3(val, ch)
	}

	// creating a new variable for sum of squares
	sum := 0

	// reading from channel
	for i := 0; i < len(slc); i++ {
		sum += <-ch
	}

	// printing sum
	fmt.Printf("Sum of squares is %d\n", sum)

	fmt.Printf("\n===================================\n\n")
}
