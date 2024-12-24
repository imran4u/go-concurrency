package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	max := 0
	start := 0

	for end, char := range s {
		if index, ok := m[char]; ok && index >= start {
			start = index + 1
		}
		clen := end - start + 1
		if max < clen {
			max = clen
		}
		m[char] = end
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Print(lengthOfLongestSubstring(s))
}
