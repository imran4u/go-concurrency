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

	fmt.Println("updating mess" + m)
	msg = m
	mutx.Unlock()
}

func main() {
	var mutx sync.Mutex

	wg.Add(2)
	go updateMessage("One", &mutx)
	// time.Sleep(100 * time.Millisecond)
	go updateMessage("Two", &mutx)
	wg.Wait()
	fmt.Println(msg)

}
