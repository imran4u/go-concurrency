package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

// in const , comma is allowed only when you are assing multiple value at once
const (
	d, e = iota, iota
	f, _
	g, _
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g)
}
