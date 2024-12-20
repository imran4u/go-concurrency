package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("Goroutine 1: Sending to ch1")
		ch1 <- 1
		ch2 <- 2

	}()

	// go func() {
	// 	fmt.Println("Goroutine 2: Sending to ch2")
	// 	// time.Sleep(3 * time.Second)
	// 	ch1 <- 2
	// 	ch2 <- 2
	// }()

	// Attempting to receive from the channels in reverse order, causing a deadlock
	fmt.Println(<-ch2) // This will block, waiting for goroutine 2 to send
	fmt.Println(<-ch1) // This will block, waiting for goroutine 1 to send

	// for {
	// 	select {
	// 	case v2 := <-ch2:
	// 		fmt.Println(v2)
	// 	case v1 := <-ch1:
	// 		fmt.Println(v1)
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("Time out")
	// 		return
	// 	}
	// }
}
