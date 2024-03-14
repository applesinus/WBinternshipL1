package main

import (
	"fmt"
	"math"
)

// creating a struct
type Point struct {
	X, Y int
}

// creating a constructor
func newPoint() Point {
	return Point{0, 0}
}

// creating a function to calculate distance using Piphagoras theorem
func distance(p1, p2 *Point) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}

func Ex24() {
	fmt.Printf("\n==========  Exercise 24:  ==========\n\n")

	// creating two points using the constructor
	p1 := newPoint()
	p2 := newPoint()
	p2.X = 4
	p2.Y = 3

	// calculating distance and printing it
	fmt.Printf("Distance: %f\n", distance(&p1, &p2))

	fmt.Printf("\n====================================\n\n")
}
