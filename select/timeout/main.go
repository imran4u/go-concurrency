package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "One"
	}()

	select {
	case m1 := <-ch1:
		fmt.Println(m1)
	case <-time.After(2 * time.Second):
		fmt.Println("Time out")
	}

	fmt.Println("End of progoram")
}
