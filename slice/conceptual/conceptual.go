package main

import "fmt"

func main() {
	var a []int
	a = append(a, 10)

	b := a
	b = append(b, 20)
	a = append(a, 30)

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}
