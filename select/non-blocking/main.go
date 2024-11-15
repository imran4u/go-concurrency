package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "One"
	}()
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		default:
			fmt.Println("Default executing")
			fmt.Println("Doing other processing ....")
			time.Sleep(2 * time.Second) // waiting more time to mock long running operation
		}
	}

	fmt.Println("End of progoram")
}
