package main

import "fmt"

func main() {
	s := "1abcddcba41abcddcba14"
	result := longestPalindromSubstring(s)
	fmt.Println(result)
}

func longestPalindromSubstring(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		temp := palindromHelper(i, i, s) // if string lenght is odd
		if len(temp) > len(result) {
			result = temp
		}
		temp = palindromHelper(i, i+1, s) // if string lenght is even
		if len(temp) > len(result) {
			result = temp
		}
	}
	return result
}

func palindromHelper(left, right int, s string) string {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
