package main

import "fmt"

func Ex11() {
	fmt.Printf("\n==========  Exercise 11:  ==========\n\n")

	// creating two example slices
	slc1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slc2 := []int{2, 4, 6, 8, 10}

	// creating a map with for storing flags
	answer := make(map[int]bool)

	// iterating through the slices and adding flags on values from the first slice
	for _, value := range slc1 {
		answer[value] = true
	}

	// iterating through the second slice and checking if the value is in the first slice using flags
	for _, value := range slc2 {
		if answer[value] {
			fmt.Printf("%d is in both slices\n", value)
		}
	}

	fmt.Printf("\n====================================\n\n")
}
