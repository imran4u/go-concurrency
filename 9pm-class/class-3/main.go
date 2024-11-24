package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	fmt.Println("starting use of channel")

	go func() {
		fmt.Println("I am in go rountine")
		time.Sleep(1 * time.Second)
		for i := 2; i < 5; i++ {
			c <- i
		}
		//c <- 2
		//c <- 3
	}()

	x := <-c // waiting
	y := <-c

	fmt.Println("x value ", x, y)

}
