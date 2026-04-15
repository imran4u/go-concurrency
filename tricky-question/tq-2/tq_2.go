package main

import "fmt"

func main() {

	var s1 []int // len-0, cap-0, prt->nil
	s2 := []int{1, 2, 3}
	n1 := copy(s1, s2) // copy(dst, src) copies min(len(dst), len(src))

	fmt.Printf("n=%d, s1=%v, s2=%v \n", n1, s1, s2) // 0, [], [1,2,3]
	fmt.Println("s1 == nil", s1 == nil)             // true

	var a []int                    // nil slice
	b := []int{}                   // empty slice
	fmt.Println("a=", a, " b=", b) // Both print [] but behave differently

}
