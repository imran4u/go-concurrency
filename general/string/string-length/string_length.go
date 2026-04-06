package main

import "fmt"

func main() {
	str := "hélló"
	count := 0
	for range str {
		count++
	}
	fmt.Println(len(str), count)
}

// Ans : 7, As string is byte array so, len([]byte) , h,l,l- 1 byte &  é ,ó - 2 byte
// so tatal = 3+2+2= 7
// range will return rune (int 32) , so it will be 5 only.
