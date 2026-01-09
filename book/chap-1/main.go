package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var data int

// 	go func() {
// 		data++
// 	}()

// 	//Note: after this also no race condition has solved still the outcome should be all three possible output, data=0,1, first 0 then 1 just before print
// 	// time.Sleep(1 * time.Millisecond) // this is bad ?
// 	if data == 0 {
// 		fmt.Printf("the value is %v. \n", data)
// 	}

// }

// // 2nd problem
func main() {
	var wg sync.WaitGroup

	for _, num := range []string{"One", "two", "three", "four", "five", "six", "seven"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num)
		}()

	}
	fmt.Println("waiting for go routine to finish")
	wg.Wait()
	fmt.Println("waiting over finished ---------")
}

//3. function closer
// func main() {
// 	var wg sync.WaitGroup

// 	for _, num := range []string{"One", "two", "three", "four", "five", "six", "seven"} {
// 		wg.Add(1)
// 		go func(num string) {
// 			defer wg.Done()
// 			fmt.Println(num)
// 		}(num)

// 	}
// 	fmt.Println("waiting for go routine to finish")
// 	wg.Wait()
// 	fmt.Println("waiting over finished ---------")
// }
