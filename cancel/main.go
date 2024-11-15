package main

import (
	"fmt"
	"runtime"
	"time"
)

func square(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for _, n := range nums {
			time.Sleep(1 * time.Second)
			select {
			case out <- n * n:
			case <-done:
				fmt.Println("canceling go routine ....")
				return
			}

		}

	}()
	return out
}

func main() {
	done := make(chan struct{}) // make a done chanel
	ch := square(done, 1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println("state A, Number of go routine ", runtime.NumGoroutine()) // here 2.

	for v := range ch {
		fmt.Println(v)
		if v == 9 {
			close(done) // it will cancel go -routine which has done channel implmentation.
		}
	}

	fmt.Println("Number of go routine ", runtime.NumGoroutine()) // it will give 1

}
