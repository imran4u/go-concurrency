package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	betterEmpty := s[:0]
	fmt.Println("better empty :", betterEmpty)
	empty := s[len(s):] // notice start from len
	fmt.Println(empty)

	// remvoe 5
	x := 5
	for i, v := range s {
		if v == x {
			s = append(s[0:i], s[i+1:]...)
		}
	}
	fmt.Println(s)
	// clear all the element of slice

	s = s[len(s):] // another way
	fmt.Println(s)

	removeAllElement()

}
func removeAllElement() {
	s := []int{1, 2, 3, 4}
	fmt.Println("removeAllElement", s)
	fmt.Println("removeAllElement len ", len(s))
	fmt.Println("removeAllElement cap", cap(s))
	s = s[0:0] // same capcity, len =0 , so reuse the same array.
	// s = s[1:1] // capacity , c-1 , len=0
	// s = s[len(s):] // cacity c-len(s), len =0

	fmt.Println("removeAllElement After empty", s)
	fmt.Println("removeAllElement len ", len(s))
	fmt.Println("removeAllElement cap", cap(s))
}
