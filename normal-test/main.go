package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// Create two decimal values, one that is zero and one that isn't
	decimal1 := decimal.NewFromFloat(0.0)
	decimal2 := decimal.NewFromFloat(0.0)
	decimal3 := decimal.NewFromFloat(5.0)
	decimal4 := decimal.NewFromFloat(-1.0)

	// Check if decimal1 and decimal2 are zero
	fmt.Println("decimal1 is zero:", decimal1.IsZero()) // true
	fmt.Println("decimal2 is zero:", decimal2.IsZero()) // true
	fmt.Println("decimal3 is zero:", decimal3.IsZero()) // false
	fmt.Println("decimal4 is zero:", decimal4.IsZero()) // false

	// Compare decimal1 and decimal2 using IsGreaterThan
	fmt.Println("decimal1 > decimal2:", decimal1.GreaterThan(decimal2)) // false (both are zero)

	// Compare decimal3 and decimal4 using IsGreaterThan
	fmt.Println("decimal3 > decimal4:", decimal3.GreaterThan(decimal4)) // true (5.0 > -1.0)

	// Comparing a zero value to a non-zero value
	fmt.Println("decimal1 > decimal3:", decimal1.GreaterThan(decimal3)) // false (0.0 is not greater than 5.0)
}
