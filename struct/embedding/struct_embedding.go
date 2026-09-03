package main

import "fmt"

type Address struct {
	City    string
	State   string
	Country string
}

type Person struct {
	Name string
	Address
	Country string
}

func main() {
	p := Person{
		Name: "imran",
		Country: "India",
		Address: Address{
			City:  "Delhi",
			State: "New Delhi",
			Country: "Bharat",
		},
	}

	fmt.Println(p.Name)
	fmt.Println(p.City)
	fmt.Println(p.State)
	fmt.Println(p.Country) //?
	fmt.Println(p.Address.Country)  // ?
}
