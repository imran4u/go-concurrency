package main

import "fmt"

// func main() {
// 	select {} // it will give dead lock;
// }

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select { // it give dead lock.
	case <-ch1:
		fmt.Println("Reading form channel 1")
	case <-ch2:
		fmt.Println("Reading form channel 2")
	}
}

/**
* IF ALL THE CASE OF SELECT IS BLOCKED THEN IT WILL PRODUCE DEADLOCK.
**/
