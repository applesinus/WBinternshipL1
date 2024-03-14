package main

import "fmt"

// creatng a power source for laptop
type power_source_220W struct {
	power int
}

// creating a power source not compatible with laptop
type power_source_360W struct {
	power int
}

// creating a laptop
type laptop struct {
	power_source power_source_220W
}

// creating an adapter to adapt power_source_360W to power_source_220W
func (power_source_360W) adapt() power_source_220W {
	return power_source_220W{220}
}

func Ex21() {
	fmt.Printf("\n==========  Exercise 21:  ==========\n\n")

	// creating a laptop with power_source_360W adapted to power_source_220W
	lenovo := laptop{power_source_360W{360}.adapt()}

	fmt.Printf("Power: %dW\n", lenovo.power_source.power)

	fmt.Printf("\n====================================\n\n")
}
