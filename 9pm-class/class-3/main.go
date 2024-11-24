package main

import (
	"fmt"
	"time"
)

func myGorountine(c chan int) {
	fmt.Println("I am in go rountine")
	time.Sleep(1 * time.Second)
	for i := 2; i < 5; i++ {
		c <- i
	}
}

func main() {

	c := make(chan int)

	fmt.Println("starting use of channel")

	go myGorountine(c)

	x := <-c // waiting
	y := <-c

	fmt.Println("x  and y value ", x, y)

}
