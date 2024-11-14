package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		fmt.Println("c value in go routine $", c)
		ch <- c
	}(1, 4)

	result, ok := <-ch

	fmt.Printf(" On Main GR got result %d, and status %t \n", result, ok)
}
