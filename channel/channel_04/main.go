// channel ownership
package main

import "fmt"

func main() {
	ch := make(chan int)
	owner := func() <-chan int {
		go func() {
			defer close(ch)

			for i := 0; i < 10; i++ {
				ch <- i
			}

		}()
		return ch
	}

	consume := func(ch <-chan int) {

		for value := range ch {
			fmt.Printf("received value = %d \n", value)
		}
		fmt.Println("Done receiving")
	}

	c := owner()
	consume(c)
}
