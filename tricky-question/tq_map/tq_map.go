package main

import "fmt"

func main() {
	var m map[string]int
	// m["a"] = 1 // panic: assignment to entry in nil map

	if v, ok := m["b"]; !ok {
		fmt.Println("no value", v)
	}
}
