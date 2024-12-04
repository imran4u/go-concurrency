package main

import (
	"fmt"
)

var count int = 5

func myGorountine(c chan int) {
	fmt.Println("I am in go myGorountine")

	// for i := 0; i < 2*count; i++ {
	// 	fmt.Println("sending value =", i)
	// 	c <- i
	// 	fmt.Println("sent... value =", i)
	// }

	close(c)
}

func main() {

	c := make(chan int, count)

	fmt.Println("starting use of channel")

	go myGorountine(c)

	for v := range c {
		fmt.Printf("value = %d \n", v)
	}

	fmt.Println("The end")

	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 4
	// c <- 5
	// close(c)

	// for v := range c {
	// 	println(v)
	// }

}
