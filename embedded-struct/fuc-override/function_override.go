package main

import "fmt"

type A struct{}

func (A) Say() {
	fmt.Println("Hello : A")
}

type B struct {
	A
}

// if this will not present then b.Say() will print "Hello : A"
func (B) Say() {
	fmt.Println("Hello : B")
}

func main() {
	b := B{}
	b.A.Say()
	//B.Say() does not override A.Say().
	// Both methods exist; B.Say() simply takes precedence when called on B.
	b.Say()
}
