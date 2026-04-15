package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // this is not a slice it is an array of fixed size 9.
	s := a[2:4]                              // l= 4-2 =2, c = 9-2 =7
	newS := append(s, 55, 66)
	fmt.Printf("len=%d, cap=%d", len(newS), cap(newS)) // 4, 7

	// another
	fmt.Println("---------")
	A := [...]int{1, 2, 3} // [3]int
	B := [3]int{1, 2, 3}   // [3]int
	fmt.Println("A==B", A == B)

	// C := [...]int{1,2,3}   // [3]int
	// D := [5]int{1,2,3}     // [5]int
	// fmt.Println("C==D", C!=D) // invalid operation: C == D (mismatched types [3]int and [5]int)c
}
