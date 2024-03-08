package main

import (
	"fmt"
	"sync"
)

// square function
func square2(x int, wg *sync.WaitGroup) {
	// printing square
	fmt.Printf("Square of %d is %d\n", x, x*x)

	// decrementing waitgroup (alias is wg.Add(-1))
	wg.Done()
}

func Ex2() {
	fmt.Printf("\n==========  Exercise 2:  ==========\n\n")

	// creating a slice with given values
	slc := []int{2, 4, 6, 8, 10}

	// creating a waitgroup with the same length as the slice to sync finish of all goroutines
	wg := new(sync.WaitGroup)
	wg.Add(len(slc))

	// creating a loop which goes through the slice
	for _, val := range slc {
		// calling square function in goroutines
		go square2(val, wg)
	}

	// waiting for all goroutines to finish
	wg.Wait()

	fmt.Printf("\n===================================\n\n")
}
