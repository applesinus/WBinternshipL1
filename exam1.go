package main

import (
	"fmt"
	"math"
)

// После сдачи подумал над возможной реализацией, и протестировал код несколькими примерами

// Почитал про slices.Concat в Go 1.22, поскольку все равно придется сортировать получившийся массив,
// я посчитал, что в данном случае эффективнее написать собственную функцию, которая гарантированно рабоатет за O(n)
func sorted_slices_merge(arr1 []int, arr2 []int) []int {
	// creating two lengths variables and a returning array
	len1, len2 := len(arr1), len(arr2)
	arr := make([]int, 0, len1+len2)

	// iterating through the two sorted slices with two indexes
	i, j := 0, 0
	for {
		if arr1[i] < arr2[j] {
			// appending lesser element and increment i
			arr = append(arr, arr1[i])
			i++
			// if arr1 ended, just appending other arr2 elements and returning
			if i == len1 {
				for ; j < len2; j++ {
					arr = append(arr, arr2[j])
				}
				return arr
			}
		} else {
			// appending lesser element and increment j
			arr = append(arr, arr2[j])
			j++
			// if arr2 ended, just appending other arr1 elements and returning
			if j == len2 {
				for ; i < len1; i++ {
					arr = append(arr, arr1[i])
				}
				return arr
			}
		}
	}
}

func exam1() {
	// creating slices
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}

	// getting the median of the two slices like they're one
	median := findMedian(arr1, arr2)
	fmt.Printf("%v\n", median)
}

func findMedian(arr1, arr2 []int) float64 {
	// merging slices
	arr := sorted_slices_merge(arr1, arr2)

	// getting the middle index and checking if there's a fractional part
	mid := float64(len(arr)) / 2
	_, fractional := math.Modf(mid)

	if fractional != 0 {
		// if there's a fractional part, the count of elements is odd
		// so just returning the middle element
		return float64(arr[int(mid)])
	} else {
		// if there's no fractional part, the count of elements is even
		// so returning the arithmetic average of the nearest elements
		return float64(arr[int(mid)-1]+arr[int(mid)]) / 2
	}
}
