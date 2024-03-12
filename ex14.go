package main

import "fmt"

func Ex14() {
	fmt.Printf("\n==========  Exercise 14:  ==========\n\n")

	// creating variables and a slice with various types
	var A int
	var B string
	var C bool
	var D chan int
	var E chan string
	slc := []interface{}{A, B, C, D, E}

	// iterating through the slice and checking the type
	for _, value := range slc {
		// switching on type
		switch value.(type) {
		case int:
			fmt.Printf("int\n")
		case string:
			fmt.Printf("string\n")
		case bool:
			fmt.Printf("bool\n")
		case chan int:
			fmt.Printf("chan int\n")
		default:
			fmt.Printf("unknown\n")
		}
	}

	fmt.Printf("\n====================================\n\n")
}
