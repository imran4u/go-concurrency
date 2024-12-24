package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"

	fmt.Printf("isPalindrome %t \n", isPalindrome(s))
	fmt.Printf("isPalindrome1 %t \n", isPalindrome1(s))
}

func isPalindrome(s string) bool {
	// remove all special character.
	// s = strings.ReplaceAll(s, " ", "")
	// s = strings.ReplaceAll(s, ",", "")
	// s = strings.ReplaceAll(s, ":", "")
	// s = strings.ToLower(s)

	//Regx to keep only alphabet and numbers.
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	// Replace all non-alphanumeric characters with an empty string
	s = re.ReplaceAllString(s, "")
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

// Alternate optimal way

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1

	for i < j {
		left := s[i]
		right := s[j]
		if isSkip(left) {
			i++
			continue
		}
		if isSkip(right) {
			j--
			continue
		}
		if left != right {
			return false
		}
		i++
		j--

	}
	return true
}

func isSkip(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'B') || (b >= '0' && b <= '9') {
		return false
	}
	return true
}
