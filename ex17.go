package main

import (
	"fmt"
	"sort"
)

func Ex17() {
	fmt.Printf("\n==========  Exercise 17:  ==========\n\n")

	// creating a slice with some values
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// searching for 7 in the slice
	index := sort.Search(len(slc), func(i int) bool { return slc[i] >= 7 })

	// old manual version
	//index := binarySearch(slc, 7)

	if index == -1 {
		fmt.Printf("7 not found\n")
	} else {
		fmt.Printf("Index of 7 is %d\n", index)
	}

	fmt.Printf("\n====================================\n\n")
}

/*
Старый код, написал его до того, как прочитал в документации, что sort.Search
	имплементирует бинарный поиск.

Логика работы алгоритма бинарного поиска (в коде будут отсылаться к этим шагам):

1. Определяем границы интервала поиска
2. Вычисляем середину интервала поиска
3. Если элемент найден, то возвращаем его индекс
4. Если элемент не найден, то уменьшаем интервал поиска
5. Поиск проходит до того, как интервал не станет состоять из нуля элементов



func binarySearch(array []int, x int) int {
	found := -1

	// stage 1
	low := 0
	high := len(array) - 1
	for low <= high {

		// stage 2
		mid := (low + high) / 2
		if array[mid] == x {
			// stage 3
			found = low + mid
			return found
		}

		// stage 4
		if array[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
		// stage 5
	}
	return found
}
*/
