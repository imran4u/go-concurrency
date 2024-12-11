package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	ch := make(chan *int, 4)
// 	array := []int{1, 2, 3, 4}
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		for _, value := range array {
// 			ch <- &value
// 		}
// 		// close(ch) // best way.

// 	}()
// 	go func() {
// 		// for value := range ch {
// 		// 	fmt.Println(*value) // what will be printed here?
// 		// }

// 		for i := 0; i < len(array); i++ {
// 			value := <-ch
// 			fmt.Println(*value) // what will be printed here?
// 		}

// 		wg.Done()
// 	}()

// 	wg.Wait()
// }

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()

	go func() {
		for value := range ch {
			fmt.Println(*value)
		}
		wg.Done()
	}()

	// New goroutine is run.
	go func() {
		var a int
		for {
			a++
		}
	}()

	wg.Wait()
}

// Note: ^ here no dead lock , because one gorountine is running infinitely.
// For deadlock all goroutine should be in waiting/blocking state
