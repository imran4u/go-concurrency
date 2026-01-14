package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	var once sync.Once
// 	var initA func()
// 	initB := func() {
// 		once.Do(initA)
// 	}

// 	initA = func() {
// 		once.Do(initB) //1.  This will not proceed until number 2 will complete.
// 	}
// 	once.Do(initA) // 2.
// }

//##  Second form of deadlock
// func main() {

// 	var once sync.Once

// 	once.Do(func() {
// 		fmt.Println("first Execusion")

// 		// This will create deadload, as another do can't start unless previous one finish.
// 		once.Do(func() {
// 			fmt.Println("second Execusion")
// 		})
// 	})

// }

// 3. No deadlock but only first once.Do will execute
func main() {

	var once sync.Once

	//1.
	once.Do(func() {
		fmt.Println("first Execusion")

	})

	//2.
	once.Do(func() {
		fmt.Println("second Execusion")
	})

}
