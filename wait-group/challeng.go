package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(m string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = m
}

func printMessage() {

	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"One",
		"two",
		"three",
		"four",
	}

	for _, w := range words {

		wg.Add(1)
		go updateMessage(w, &wg)
		wg.Wait()
		printMessage()
	}

	fmt.Println("End of main go routine after all go routine")

}
