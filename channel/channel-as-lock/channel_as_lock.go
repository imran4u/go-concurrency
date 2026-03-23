package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	count := 0
// 	wg := sync.WaitGroup{}

// 	for range 1000 {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			count++
// 		}()

// 	}

// 	wg.Wait()

// 	fmt.Println("count=", count)
// }

/**
Above has race coudition it can solve by either way
1. useing mutext lock
2. using atomic eg. atomic.AddInt32(&count, 1)
3. using channel as lock.
**/

//using chnnel as lock.

func main() {
	count := 0
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 1) // buffered channel with 1 value. 

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- struct{}{}
			count++
			<-ch
		}()

	}

	wg.Wait()

	fmt.Println("count=", count)
}
