package main

import (
	"fmt"

	"github.com/imran4u/go-concurrency/general/stack"
)

func main() {
	s := stack.Stack{}
	fmt.Println(s)
	fmt.Println("isEmpty", s.IsEmpty())
	s.Push(2)
	fmt.Println(s.Peek())

	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())

}
