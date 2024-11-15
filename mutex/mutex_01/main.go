package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var balance int

func main() {
	// without all process you may get correct result without mutex,
	//so use all 4 cpu i.e pallel goroutine
	runtime.GOMAXPROCS(4)

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}
	widthDraw := func(amount int) {
		defer mu.Unlock()
		mu.Lock()
		balance -= amount
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			deposit(2)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			widthDraw(2)
		}
	}()
	wg.Wait()
	fmt.Println("Balace is ", balance)

}
