package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4) // run parallel program

	var wg sync.WaitGroup
	var count int64

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {

				//count++ // it will give wrong result

				atomic.AddInt64(&count, 1) // now result will correct 50*1000 = 50,000
			}

		}()
	}
	wg.Wait()
	fmt.Println("count value now ", count)
}
