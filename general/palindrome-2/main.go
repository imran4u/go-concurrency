package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"

	fmt.Printf("isPalindrome %t \n", isPalindrome(s))
}

func isPalindrome(s string) bool {
	// remove all special character.
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ToLower(s)

	fmt.Println(s)
	l := len(s)
	i := 0
	j := l - 1

	for i <= j {

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
