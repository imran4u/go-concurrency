package main

import (
	"fmt"
	"strings"
)

// 0, 1,1,2,3,5,8,13,21,34 ....
func main() {
	max := 30
	for n := 0; n < max; n++ {
		fmt.Printf("for number = %d, fibonaci= %d \n", n, FibonacciNumber(n))
	}
	var sb strings.Builder
	for n := 0; n < max; n++ {
		sb.WriteString(fmt.Sprintf("%d,", FibonacciNumber(n)))

	}
	fmt.Println(sb.String())
}

func FibonacciNumber(n int) int {
	if n < 2 {
		return n
	}
	// s := make([]int, 0, 10) // no make sure to give size and capacity, first is size and then capacity
	// // s := []int{} // using slice
	// s = append(s, 0)
	// s = append(s, 1)

	// for i := 2; i <= n; i++ {
	// 	s = append(s, s[i-1]+s[i-2])
	// }

	//Array way
	s := make([]int, n+1) // create slice of capacity n+1
	s[0] = 0
	s[1] = 1

	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]

}
