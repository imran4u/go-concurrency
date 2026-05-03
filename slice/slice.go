package main

import "fmt"

func main() {
	s := make([]int, 10)
	// s := new([]int) // it is quivalent to -> var s []int
	// so you must have to initia
	// *s = []int{1,2,3}
	// or
	// *s = make([]int, 0)

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println("after 10 slice : ", s)
	fmt.Println("after 10 size : ", len(s))
	fmt.Println("after 10 cap : ", cap(s))

	// with new

	sp := new([]int) // it is quivalent to -> var s []int
	// so you must have to initia
	// *s = []int{1,2,3}
	// or
	// *sp = make([]int, 0)

	fmt.Println("inital slice : ", sp)
	fmt.Println("inital size : ", len(*sp))
	fmt.Println("inital cap : ", cap(*sp))

}
