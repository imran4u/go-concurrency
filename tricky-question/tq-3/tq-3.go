package main

import "fmt"

func service() {
	fmt.Println("Hello from service")
}

func main() {
	fmt.Println("main started")
	go service()

	select {} // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Main stoped")

}
