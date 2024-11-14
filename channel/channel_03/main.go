// Channel direction
package main

import "fmt"

// Chanel one is only write , not for read, Read operation will give compilation error.
func genMessage(ch chan<- string) {
	ch <- "Message"
}

// Read from channel one and write to channel two
func relayMessage(in <-chan string, out chan<- string) {
	value := <-in
	out <- value
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go genMessage(ch1)
	go relayMessage(ch1, ch2)
	v := <-ch2
	fmt.Println(v)
}
