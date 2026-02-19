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

	var counter int16
	timeout := time.After(3 * time.Second) // In production always use channel out side fo slect case.

	// Note: seclect run all case and check which one is ready to execute and <-time.After(1 * time.Second) return false and will go default.

	for {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		// case <-time.After(1 * time.Second): // Note on each loop new timer is going to created. if this value is more than default sleep, it will never execute.
		case <-timeout:
			fmt.Println("Time out")
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("I am default")

		}
		if counter == 10 {
			break
		}
		counter++
	}

	fmt.Println("End of progoram")
}
