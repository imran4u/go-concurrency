package main

import (
	"fmt"
	"sync"
	"time"
)

// 7 % 4 = 3
//0,2,4,6,8, even

//slect { }

func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)
	done := make(chan bool)
	wg.Add(1)
	go consumer(ch1, ch2, done)

	go func() {
		for i := 0; i < 20; i++ {
			if i%2 == 0 {
				ch2 <- i
			} else {
				ch1 <- i
			}
			time.Sleep(300 * time.Millisecond)
		}

		close(ch1)
		close(ch2) // Close channels when done sending values
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Work done")

}

func consumer(ch1 chan int, ch2 chan int, done chan bool) {
	for {
		select {
		case v := <-ch1:
			fmt.Println(">>>>>>>>>>>>>. odd chan 1 :", v)
		case v := <-ch2:
			fmt.Println("even chan 2 :", v)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("############### Time expired")
			done <- true
			return
		}
	}
}
