package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var onece sync.Once

	load := func() {
		fmt.Println("Loading should be only onece")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// load() // this will 10 time from all go routine.
			onece.Do(load) // it will call only once insted from multipul go routine call

		}()
	}
	wg.Wait()
}
