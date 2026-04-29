package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
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

}
