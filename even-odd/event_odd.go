package main

import (
	"fmt"
)

func main() {
	even := []int{0, 2, 4, 6, 8}
	odd := []int{1, 3, 5, 7, 9}
	ch := make(chan int)
	evenCh := make(chan struct{})
	oddCh := make(chan struct{})

	go func() {

		for i, v := range even {
			<-evenCh
			ch <- v

			// only signal if not last, to avoid go routine leak.
			if i != len(odd)-1 {
				evenCh <- struct{}{}
			}
		}

	}()

	go func() {

		for _, v := range odd {
			<-oddCh
			ch <- v
			evenCh <- struct{}{}
		}

	}()

	//start with even
	evenCh <- struct{}{}

	count := len(even) + len(odd)
	for range count {
		fmt.Println(<-ch)
	}
}
