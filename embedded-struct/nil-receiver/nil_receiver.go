package main

import (
	"fmt"
)

type A struct{}

func (a *A) Say() {
	fmt.Println("A", a)
}

func main() {
	var a A
	fmt.Printf("a=%v\n", a)
	fmt.Println("a==nil", &a == nil) //Since a exists, its address cannot be nil

// 	👉 Rule:
// Only pointers, maps, slices, channels, functions, and interfaces can be nil — not struct values

	a.Say()

	//
	var b *A= nil

	fmt.Println("b==nil",b==nil)
	b.Say()

}
