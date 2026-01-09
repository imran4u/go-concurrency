package main

import (
	"fmt"
	"strings"
)

func main() {
	var address string
	// String literal
	address = "String Literal need scape charater for multiline \nBanglore karnataka India"
	fmt.Println(address)

	// Raw string
	address = `Raw string multi line support 
	Banglore
	karnataka
	India`
	fmt.Println(address)

	test := `A test string`

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
	fmt.Println(`>>>Rune print`)
	fmt.Println('A') // 65
	fmt.Println('a') // 97

	// Reverse the number
	number := "123456789"
	size := len(number)
	var sb strings.Builder
	for i := size - 1; i >= 0; i-- {
		sb.WriteByte(number[i])
	}
	fmt.Println(" number = ", number)
	fmt.Println("reversed number = ", sb.String())

	// byte vs rune
	// Array of bytes (ASCII characters) ASCII support only  128 character inclusing english letter
	bn := []byte{'A', 'B', 'C', 'D'}

	// Accessing element at index 1
	fmt.Println(bn[1]) // Output: 66 (ASCII value of 'B')

	// Unicode : cover all written languages in the world and other symbols
	// Array of runes (Unicode characters)
	rn := []rune{'A', 'B', 'C', 'D'}
	// Accessing element at index 1
	fmt.Println(rn[1]) // Output: 66 (Unicode value of 'B')

}
