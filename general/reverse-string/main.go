package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	size := len(s)
	var sb strings.Builder
	for i := size - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	fmt.Println(s)
	fmt.Println(sb.String())
}
