package main

import "fmt"

func main(){
	a := []string{"a", "b", "c"}
	// a = nil // a-> nil so, a->[], len ->0, cap ->0

	a = a[:0] // a->[], len ->0, cap ->3
	fmt.Println(a, len(a), cap(a))
}
