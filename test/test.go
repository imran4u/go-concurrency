package main

import "fmt"

func main() {
	r := fmt.Sprintf("Hello friend %d", "three") // go vet will catch this code 
	fmt.Println(r)
}
