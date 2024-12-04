package main

import (
	"fmt"
	"sync"
)

var message string

func updateMessage(msg string, mutex *sync.Mutex) {
	mutex.Lock()
	message = msg
	mutex.Unlock()
}

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(20)

	updateMessage("I am from main", &mutex)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			updateMessage(fmt.Sprintf("I am go routine with value = %d", i), &mutex)
		}()
	}

	wg.Wait()

	fmt.Println(message)
}
