package main

import "fmt"

func Ex8() {
	fmt.Printf("\n==========  Exercise 8:  ==========\n\n")

	// getting number, bit position and value
	fmt.Printf("Enter the number: ")
	var n int64
	fmt.Scan(&n)
	fmt.Printf("Enter the bit position: ")
	var bit int
	fmt.Scan(&bit)
	fmt.Printf("Enter the value of bit (0 or 1): ")
	var value int
	fmt.Scan(&value)

	if value == 0 {
		// if it's 0, use operator BITCLEAR (AND NOT) to set any value to 0
		// here is a trick to set bit to 0 using bitclear operation
		// basically when you bitclear with 1, you get 0 anyways
		n = n &^ (1 << bit)
	} else {
		// if it's 1, use operator OR to set any value to 1
		n = n | (1 << bit)
	}
	// I think this is the most efficient way to set bit to 0 or 1

	fmt.Printf("\n===================================\n\n")
}
