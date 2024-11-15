package main

import (
	"sync"
)

// Create Pipe line of generate() ->square() -> print()
// each state will be go routine and conneced with channel.
// let assume square is hig computation so break it and merge in
// an achieve fan-in fan-out.

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

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// write on new channel out
	output := func(c <-chan int) {
		defer wg.Done()
		for cv := range c {
			out <- cv
		}

	}

	//iterate on channle
	for _, cv := range cs {
		go output(cv)
	}

	// wait and close channel if all done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	in := generate(5, 7, 8)

	//fan-out , distribute
	ch2 := square(in)
	ch3 := square(in)

	ch := merge(ch2, ch3) // Note: here order of result might be change from input that is expected.

	for value := range ch {
		println(value)
	}

}
