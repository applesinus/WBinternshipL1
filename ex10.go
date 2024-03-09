package main

import "fmt"

func Ex10() {
	fmt.Printf("\n==========  Exercise 10:  ==========\n\n")

	// creating a map for output and a slice with given values
	temps := make(map[int]([]float32))
	given_vars := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// iterating through the slice and adding values to the map
	for _, value := range given_vars {
		// int(value) / 10 * 10 - rounding to the nearest 10
		temps[int(value)/10*10] = append(temps[int(value)/10*10], value)
	}

	fmt.Printf("Temperatures: %v\n", temps)

	fmt.Printf("\n====================================\n\n")
}
