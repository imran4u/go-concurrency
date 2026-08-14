package main

import "fmt"

/*
output :
	Integer result (2 + 3): 5
	Float result (2.5 + 3.1): 5.6
	String result (Hello + Programmer): Hello Programmer
*/

func main() {
	resultInt := add(2, 3)
	fmt.Println("Integer result (2 + 3):", resultInt)

	resultFloat := add(2.5, 3.1)
	fmt.Println("Float result (2.5 + 3.1):", resultFloat)

	resultStr := add("Hello ", "Programmer")
	fmt.Println("String result (Hello + Programmer):", resultStr)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

// func add(a, b int) int {
// 	return a + b
// }
