package main

import (
	"fmt"
	"strings"
)

/*

=====   ПРИМЕР КОДА   =====

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


===== ОТВЕТ =====

1. Использование глобальной переменной justString - возможная причина
	непредвиденного поведения программы, учитывая, что при ее
	использовании нет лока мьютекса

2. Использование операции (1 << 10) вместо константы 1024 - лишняя операция,
	которая к тому же уменьшает читаемость кода

3. Просто догадка, но хранить в памяти такую большую строку разве эффективно?

4. Странно вообще создавать переменную v, она используется только один раз,
	можно объединить две строки и сделать эту область памяти неименованной.
	Если только для читаемости кода.

*/

// example of a huge string creator function
func createHugeString(size int) string {
	// creating a string builder with given size
	var builder strings.Builder
	builder.Grow(size)

	// filling the string builder with 'a'
	for i := 0; i < size; i++ {
		builder.WriteRune('a')
	}

	// returning the string from the builder
	return builder.String()
}

// refactored function
func someFunc() string {
	// creating a huge string by calling the function
	justString := createHugeString(1024)[:100]
	return justString
}

func Ex15() {
	fmt.Printf("\n==========  Exercise 15:  ==========\n\n")

	// calling the function
	fmt.Printf("justString: %v\n", someFunc())

	fmt.Printf("\n====================================\n\n")
}
