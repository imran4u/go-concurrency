package main

import (
	"fmt"
	"sync"
)

func main() {
	input := "aacabdkacaa"
	var longest string

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= len(input); i++ {
		wg.Add(1)
		go checkPalindrome(input, i, ch, &wg)
	}

	// important to close the channel, if you are using range.
	go func() {
		wg.Wait()
		close(ch)
	}()

	for chVal := range ch {
		if len(chVal) > len(longest) {
			longest = chVal
		}
	}

	fmt.Println(longest)
}

func checkPalindrome(input string, window int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i+window <= len(input); i++ {
		sub := input[i : i+window]
		if isPalindrome(sub) {
			ch <- sub
			return
		}
	}
	ch <- ""
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
