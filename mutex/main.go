package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(m string, mutx *sync.Mutex) {
	defer wg.Done()
	mutx.Lock()
	msg = m
	mutx.Unlock()
}

func main() {
	var mutx sync.Mutex

	wg.Add(2)
	go updateMessage("One", &mutx)
	go updateMessage("One", &mutx)
	wg.Wait()
	fmt.Println(msg)

}
