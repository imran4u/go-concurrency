package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func tea() {
	defer wg.Done()
	fmt.Println("tea")
}

func bread() {
	time.Sleep(10 * time.Microsecond)
	fmt.Println("Bread")
	wg.Done()
}
func main() {

	wg.Add(4)

	go tea()
	go bread()

	go func() {
		fmt.Println("Ananomous fuction")
		wg.Done()
	}()

	assign := func() {
		defer wg.Done()

		fmt.Println("Assign")
	}
	go assign()

	wg.Wait()

}
