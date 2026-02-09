package main

import "fmt"

func main() {

	input := []int{}
	input = append(input, 1)

	fmt.Println("Hello slice")
	fmt.Printf("input length =%d \n", len(input))
	fmt.Printf("input capacity =%d \n", cap(input))
	fmt.Printf("access out of inex = %d \n", input[10]) // index out of bound

	///
	s := make([]int, 3, 5)

	fmt.Printf("s length =%d \n", len(s))
	fmt.Printf("s capacity =%d \n", cap(s))
}
