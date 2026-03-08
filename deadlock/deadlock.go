package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Println("Before sending")
	ch <- 4
	fmt.Println("After sending")
}
