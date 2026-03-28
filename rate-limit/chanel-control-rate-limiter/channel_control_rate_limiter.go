package main

import (
	"fmt"
	"sync"
	"time"
)

//Normal channel
/**
func main() {
	var wg sync.WaitGroup
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-tick.C
			fmt.Println("Allowed", i)

		}(i)
	}
	wg.Wait()
}
**/

// Buffered channel

func main() {
	wg := sync.WaitGroup{}
	token := make(chan struct{}, 2) // bucket capacity 2

	// refil bucked at rate of 500 ms, so it will allow only two call every sec
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case token <- struct{}{}:
			default:
				// do nothing
			}
		}

	}()

	// 10 concurent call
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// <-ch
			// fmt.Println("Allowed", i)

			if isAllowed(token) {
				fmt.Println("Allowed", i)
			} else {
				fmt.Println("NotAllowed", i)
			}

		}(i)
		time.Sleep(300 * time.Millisecond)
	}
	wg.Wait()

}

func isAllowed(tokens <-chan struct{}) bool {
	select {
	case <-tokens:
		return true
	default:
		return false
	}
}
