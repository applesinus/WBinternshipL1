package main

import "fmt"

/*

Это могло быть задание на арифметику больших чисел,
но 2^20*2^20 имеет 13 знаков, а int64 - 18 знаков
причем на любой архитектуре.

Если это задание на знание типов данных - то ок, если
на длинную арифметику - то по тз это не очевидно.

*/

func Ex22() {
	fmt.Printf("\n==========  Exercise 22:  ==========\n\n")

	// getting a and b
	var a, b int64
	fmt.Printf("Enter a and b: ")
	fmt.Scan(&a, &b)

	// using standart math operators
	fmt.Printf("a*b = %d\n", a*b)
	fmt.Printf("a/b = %d\n", a/b)
	fmt.Printf("a+b = %d\n", a+b)
	fmt.Printf("a-b = %d\n", a-b)

	fmt.Printf("\n====================================\n\n")
}
