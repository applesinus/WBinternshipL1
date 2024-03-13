package main

import (
	"fmt"
	"strings"
)

func Ex19() {
	fmt.Printf("\n==========  Exercise 19:  ==========\n\n")

	// getting a string
	var str string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)

	// creating a string builder
	var builder strings.Builder
	builder.Grow(len(str))

	// transforming the string to a slice of runes
	runes := []rune(str)

	// iterating through the slice in reverse order and writing each rune to the builder
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	fmt.Printf("Reversed string: %s\n", builder.String())

	fmt.Printf("\n====================================\n\n")
}
