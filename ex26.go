package main

import "fmt"

func Ex26() {
	fmt.Printf("\n==========  Exercise 26:  ==========\n\n")

	// getting a string
	fmt.Printf("Enter the string: ")
	var str string
	fmt.Scan(&str)

	// creating a map of runes and their counts
	chars := make(map[rune]int)

	// iterating through the string
	for _, c := range str {
		chars[c] = chars[c] + 1
		// checking if the char is in the map (so it's a duplicate)
		if chars[c] > 1 {
			// if it's the case, print false and return
			fmt.Printf("False\n")
			fmt.Printf("\n====================================\n\n")
			return
		}
	}

	// if the string is unique, print true
	fmt.Printf("True\n")
	fmt.Printf("\n====================================\n\n")
}
