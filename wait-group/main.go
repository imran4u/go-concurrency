package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // decrement wg counter by 1
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"One",
		"two",
		"three",
		"four",
	}

	wg.Add(len(words)) // increment the wg counter by length of words

	for _, w := range words {
		go printSomething(w, &wg)
	}
	wg.Wait()
	wg.Add(1) // add one more because I was to print one more final stegment in go rountine
	go printSomething("I am the last go routine", &wg)
	wg.Wait()

	fmt.Println("End of main go routine after all go routine")

}
