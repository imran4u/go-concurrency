package main

import (
	"fmt"
)

// func main() {
// 	ch := make(chan int)

// 	go func(a, b int) {
// 		c := a + b
// 		fmt.Println("c value in go routine $", c)
// 		ch <- c
// 	}(1, 4)

// 	result, ok := <-ch

// 	fmt.Printf(" On Main GR got result %d, and status %t \n", result, ok)
// }

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		fmt.Println("c value in go routine $", c)
		ch <- c
		ch <- 3
		ch <- 9
		close(ch)
	}(1, 4)

	result, ok := <-ch

	fmt.Printf(" On Main GR got result %d, and status %t \n", result, ok)

	for cv := range ch { // range loop will continue unless close(ch) called.
		fmt.Printf(" cv value %d \n", cv)
	}
}
