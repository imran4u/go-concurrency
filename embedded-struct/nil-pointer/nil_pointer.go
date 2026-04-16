package main

import "fmt"

type A struct{}

func (a *A) Hello() {
	fmt.Println("Hello")
}

type B struct {
	*A
}

func main() {
	b := B{}
	fmt.Println("b.A == nil", b.A == nil)
	b.Hello() // Even though A is embedded as *A, its methods are promoted to B.
}
