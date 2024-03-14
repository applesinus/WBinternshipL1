package main

import "fmt"

func Ex23() {
	fmt.Printf("\n==========  Exercise 23:  ==========\n\n")

	// creating a slice
	slc := []int{1, 2, 3, 4, 5}

	// deleting an index to delete
	fmt.Printf("Enter an index to delete: ")
	var index int
	fmt.Scan(&index)

	// deleting an item
	slc = append(slc[:index], slc[index+1:]...)
	fmt.Printf("New slice is: %v\n", slc)

	fmt.Printf("\n====================================\n\n")
}
