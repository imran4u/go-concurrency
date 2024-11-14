package main

import "fmt"

func main() {
	// Note: here both buffered and unbufferd both will work.
	// ch := make(chan int) // un buffered channel
	size := 10
	ch := make(chan int, size)

	go func() {
		for i := 0; i < size; i++ {
			ch <- i
		}
		// chanel must be close after use as reading is used range operator
		close(ch)
	}()

	// loop will exit on closed channel
	// on range chanel will return value as in FIFO fashion
	for value := range ch {

		fmt.Println(value)
	}

}
