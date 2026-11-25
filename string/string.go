package main

import "fmt"

func main() {
	s := "hé"
	rs := []rune(s)
	fmt.Println(len(s))
	fmt.Println(len(rs))
}
