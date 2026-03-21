package main

import "fmt"

// func main() {

// 	s1 := make([]int, 0, 3)   // [x,x,x]-cap, len=0
// 	s2 := append(s1, 1, 2, 3) // [1,2,3]
// 	s3 := append(s2, 4, 5) // out of capacity so new slice of [1,2,3,4,5]
// 	s2[0] = 10  // [10,2,3]

// 	fmt.Println(s1) //  []
// 	fmt.Println(s2) //[10 2 3]
// 	fmt.Println(s3) // [1 2 3 4 5]
// }

/**
out put
[]
[10 2 3]
[1 2 3 4 5]
**/

func main() {

	s1 := make([]int, 0, 8) // 3-> 8 // len=0, cap : [x,x,x,x,x,x,x,x]
	s2 := append(s1, 1, 2, 3) // [1,2,3,x,x,x,x,x]
	s3 := append(s2, 4, 5) // in capacity so, same array [1,2,3,4,5,x,x,x]
	s2[0] = 10  // [10,2,3,4,5,x,x,x]
	
	fmt.Println(s1) // []
	fmt.Println(s2) //[10,2,3]
	fmt.Println(s3) //[10,2,3,4,5]
}
/**

[]
[10 2 3]
[10,2,3,4,5]
**/