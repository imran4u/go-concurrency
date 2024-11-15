package main

import "fmt"

// Create Pipe line of generate() ->square() -> print()
// each state will be go routine and conneced with channel.

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {

	// ch1 := generate(5, 7, 8)
	// ch2 := square(ch1)

	// for value := range ch2 {
	// 	println(value)
	// }
	// Above is also fine but lets make it short

	for value := range square(generate(2, 3, 4, 5)) {
		fmt.Println("square value =", value)

	}

}
