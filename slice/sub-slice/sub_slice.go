package main

import "fmt"

/**
Note: s := make([]int, 0, 1000)
s1:= s[a:b] // len = b-a , cap = cap - a
s1:= s[a:b:c] // len = b-a , cap = c-a  , always 0<=a<=b<=c

**/

func subSlice(s []int) []int {
	return s[0:2]
}

// func main() {
// 	s := make([]int, 0, 1000)

// 	for i :=range 1000 {
// 		s= append(s, i)
// 	}
// 	s1 := subSlice(s)

// 	fmt.Printf("s len =%d, cap = %d \n", len(s), cap(s))
// 	fmt.Printf("s1 len %d, cap = %d", len(s1), cap(s1))
// }
/**
s len =1000, cap = 1000
s1 len 2, cap = 1000
**/

func main() {
	s := make([]int, 0, 1000)

	// s1 := subSlice(s)
	s1 := s[0:2]
	fmt.Printf("s len =%d, cap = %d \n", len(s), cap(s))
	fmt.Printf("s1 len %d, cap = %d \n", len(s1), cap(s1))
	fmt.Println("s=", s)
	fmt.Println("s1=", s1)
	s1[0] = 10
	fmt.Println("after s=", s) // Even though underlying array changed, s cannot "see" it.Because len(s) = 0 : Slice controls visibility via length, not memory content.
	fmt.Println("after s1=", s1)
	s = append(s, 5)

	fmt.Println("after s update, s=", s) 
	fmt.Println("after s update, s1=", s1)
}

/**


s len =0, cap = 1000 
s1 len 2, cap = 1000 
s= []
s1= [0 0]
after s= []
after s1= [10 0]
after s update, s= [5]
after s update, s1= [5 0]
**/
