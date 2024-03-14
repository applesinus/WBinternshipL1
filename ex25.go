package main

import (
	"fmt"
	"time"
)

// creating a sleep function
func my_sleep(sec int) {
	// getting the current time
	start := time.Now()

	// iterating until the time passed
	for {
		if time.Since(start).Seconds() > float64(sec) {
			fmt.Printf("\nRING! RING! RING!\n%d seconds passed\n", sec)
			break
		}
	}
}

func Ex25() {
	fmt.Printf("\n==========  Exercise 25:  ==========\n\n")

	// getting a sleep time
	fmt.Printf("Set a sleep time in seconds: ")
	var sec int
	fmt.Scan(&sec)

	// calling the sleep function
	my_sleep(sec)

	fmt.Printf("\n====================================\n\n")
}
