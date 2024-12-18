package main

import "fmt"

func safeFunction() {
	defer func() {
		if r := recover(); r != nil { // this will capture panic and recover to terminate program.
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("No Recovered from panic:")
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // This will not be executed
}

func main() {
	safeFunction() // Panic is recovered here
	fmt.Println("After recovery")
}
