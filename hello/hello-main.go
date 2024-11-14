package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	// Direct call without go routine
	fun("direct call")

	// goroutine function call
	go fun("goroutine-1")

	// goroutine with anonymous function
	go func() {
		fun("anonymous ")
	}()

	// goroutine with function value call
	fv := fun
	go fv("function value call")

	// wait for goroutines to end time.Sleep is not a good idea
	//but we are using for sake of similicity as of now.
	fmt.Println("waiting to complete go routine")
	time.Sleep(5 * time.Millisecond)

	fmt.Println("done..")
}
