package main

import "fmt"

func main() {
	var m map[string]string //nil map
	fmt.Println(m)

	r := m["ali"]
	fmt.Println("r value =", r)
	m["ali"] = "update" //nil map don't allow to write though read is permitted

	ru := m["ali"]
	fmt.Println("updated ru value =", ru)

}
