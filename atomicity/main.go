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
	var couter2 atomic.Uint64 // new way in go 1.9
	var b atomic.Bool

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {

				//count++ // it will give wrong result

				atomic.AddInt64(&count, 1) // now result will correct 50*1000 = 50,000
				couter2.Add(1)
				b.Store(true)
			}

		}()
	}
	wg.Wait()
	fmt.Println("count value now ", count)

	fmt.Println("count value from counter2 :  ", couter2.Load())

	fmt.Println("Boolean value :  ", b.Load())
}
