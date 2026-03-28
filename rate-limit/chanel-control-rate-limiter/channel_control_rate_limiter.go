package main

import (
	"fmt"
	"sync"
	"time"
)

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
