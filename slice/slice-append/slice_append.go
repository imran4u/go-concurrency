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

// func main() {

// 	s1 := make([]int, 0, 8) // 3-> 8 // len=0, cap : [x,x,x,x,x,x,x,x]
// 	s2 := append(s1, 1, 2, 3) // [1,2,3,x,x,x,x,x]
// 	s3 := append(s2, 4, 5) // in capacity so, same array [1,2,3,4,5,x,x,x]
// 	s2[0] = 10  // [10,2,3,4,5,x,x,x]

// 	fmt.Println(s1) // []
// 	fmt.Println(s2) //[10,2,3]
// 	fmt.Println(s3) //[10,2,3,4,5]
// }
/**

[]
[10 2 3]
[10,2,3,4,5]
**/

// func main() {

// 	s := make([]int, 0, 3) // s=[x,x,x]

// 	s1 := append(s, 1) // s1= [1,x,x]
// 	s2 := append(s, 2)	// s2= [2,x,x]

// 	fmt.Println("s =", s) // len =0 so, []
// 	fmt.Println("s1 =", s1) // len=1, so last update value [2]
// 	fmt.Println("s2 =", s2) // len=1, so last update value [2]

// }

/**

s = []
s1 = [2]
s2 = [2]
*/

// vvI
// func main() {
// 	s := []int{1, 2, 3} // s=[1,2,3]

// 	a := append(s[:1], 100) // s=[1] & append 100 so a= [1,100], s= [1,100,2]
// 	b := append(s[:1], 200) // b = [1,200] , s = [1,200,2]

// 	fmt.Println("a=", a) // now see the s with len=2 [1,200]
// 	fmt.Println("b=", b) // now see the s with len=2 [1,200]
// 	fmt.Println("s=", s) // now see the final s  [1,200,2]
// }

// Good example

func main() {

	s := []int{1, 2, 3, 4}

	a := s[:2]         // [1 2]
	b := append(a, 10) // ? [1,2,10,4]  but len = 3
	c := append(a, 20) // ? [1,2,20,4]  but len = 3

	fmt.Println("s =", s) // [1,2,20,4]
	fmt.Println("a =", a) //[1,2,20,4]   but len 2
	fmt.Println("b =", b) //[1,2,20,4]   but len 3
	fmt.Println("c =", c) //[1,2,20,4]   but len 3
}

/**

s = [1 2 20 4]
a = [1 2]
b = [1 2 20]
c = [1 2 20]
**/
