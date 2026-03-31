package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	mut := sync.Mutex{}
	c := sync.NewCond(&mut)

	var mapRes = make(map[string]string)
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(mapRes) == 0 {
			fmt.Println("length is zero")
			c.Wait()
			fmt.Println("Wait over signal")
		}

		fmt.Println(mapRes["first"])
		c.L.Unlock()
	}()

	//main go routine
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		c.L.Lock()
		mapRes["first"] = "First value"
		c.Signal()
		c.L.Unlock()
	}()

	wg.Wait()
	fmt.Println("Done all work")

}
