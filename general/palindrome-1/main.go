package main

import "fmt"

func main() {
	result := isPalindrome(10)
	fmt.Printf("isPalendrome : %t \n", result)
}

// this is not efficient way, as we need to change the int to string.
// func isPalindrome(x int) bool {

// 	if x < 0 {
// 		return false
// 	}
// 	// s := fmt.SPrintf("%d", x)
// 	s := fmt.Sprintf("%d", x)
// 	i := 0
// 	j := len(s) - 1
// 	for i <= j {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

// >>>. Efficient way.
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { // either -ve or ends with 0 except 0 itself.
		return false
	}
	temp := x
	rever := 0
	for temp > 0 {
		rever = rever*10 + temp%10
		temp = temp / 10
	}

	return rever == x
}
