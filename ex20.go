package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ex20() {
	fmt.Printf("\n==========  Exercise 20:  ==========\n\n")

	// getting a string using bufio to read string with spaces from stdin
	var str string
	fmt.Printf("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		// catching errors
		fmt.Printf("Error: %v\n", err)
	}

	// creating a string builder
	var builder strings.Builder
	builder.Grow(len(str))

	// transforming the string to a slice of runes
	str = strings.TrimSpace(str)
	words := strings.Fields(str)

	// iterating through the slice in reverse order and writing each rune to the builder
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		builder.WriteRune(' ')
	}

	fmt.Printf("Reversed order of words: %s\n", builder.String())

	fmt.Printf("\n====================================\n\n")
}
