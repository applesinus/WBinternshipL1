package main

import "fmt"

func Ex12() {
	fmt.Printf("\n==========  Exercise 12:  ==========\n\n")

	// creating a slice with given values
	given_strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// creating a set for output
	set := make([]string, 0)

	// iterating through the slice
	for _, value := range given_strings {
		flag := true
		for _, val := range set {
			// checking if the value is already in the set
			if value == val {
				flag = false
			}
		}

		// if the value is not in the set, add it
		if flag {
			set = append(set, value)
		}
	}

	// printing the set
	fmt.Printf("Set: %v\n", set)

	fmt.Printf("\n====================================\n\n")
}
