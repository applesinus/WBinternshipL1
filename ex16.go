package main

import (
	"fmt"
	"sort"
)

func Ex16() {
	fmt.Printf("\n==========  Exercise 16:  ==========\n\n")

	// creating a slice with some values
	slc := []int{10, 9, 5, 6, 8, 7, 3, 4, 2, 1}

	// old manual version
	// slc = quickSort(slc, 0, len(slc)-1)

	// sorting the slice
	sort.Slice(slc, func(i, j int) bool { return slc[j] > slc[i] })
	fmt.Printf("Sorted array is: %v\n", slc)

	fmt.Printf("\n====================================\n\n")
}

/*
Старый код, написал его до того, как прочитал в документации, что sort.Slice имплементирует квиксорт.

Логика работы квиксорта (в коде комменты будут отсылаться к этим шагам):

0. Если размер массива больше 1, то начинаем/продолжаем сотировку
1. Определение опорного элемента (в данной вариации - последний элемент) массива
2. Проход по массиву с двумя индексами: i, j
3. Если текущий элемент с индексом j меньше опорного элемента, то меняем элементы
	с индексами i, j местами, i увеличивается на единицу
4. j увеличивается на единицу
5. Когда j дошел до конца массива, меняем местами опорный элемент и элемент с индексом i,
	при этом он встает на свое место, то есть все элементы левее - меньше, а правее - больше
6. Подаем рекурсивно на сортировку левую и правую части массива, при этом
	опорный элемент не включается ни в одну из них

Отдельно отмечу, что квиксорт можно распараллелить, но тогда придется выделять больше памяти,
поскольку либо каждый поток будет работать с отдельным подмассивом, либо придется использовать
мьютексы, а тогда в этом не будет смысла. Но в тз не было сказано ни про параллельность, ни про
предпочтения по оптимизации с точки зрения памяти/скорости/производительности.



// partition function
func partition(arr []int, low, high int) ([]int, int) {
	// stage 1
	pivot := arr[high]

	// stage 2
	i := low
	for j := low; j < high; j++ {
		// stage 3
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		// stage 4 (j++ is in the declaration of for)
	}
	// stage 5
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

// quicksort function
func quickSort(arr []int, low, high int) []int {
	// stage 0
	if low < high {
		var p int
		arr, p = partition(arr, low, high)

		// stage 6
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
*/
