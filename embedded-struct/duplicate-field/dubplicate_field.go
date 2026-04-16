package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

type D struct {
	A
}

func main() {
	c := C{A: A{Name: "Ali"}, B: B{Name: "BName"}}
	fmt.Println(c.A.Name)
	fmt.Println(c.B.Name)
	//Ambiguity only happens when multiple promoted fields share the same name at the same depth.
	// fmt.Println(c.Name) // ambiguous selector

	d := D{A: A{Name: "aName"}}
	fmt.Println(d.Name) // field promotion
}
