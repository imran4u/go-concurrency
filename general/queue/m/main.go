package main

import (
	"fmt"

	"github.com/imran4u/go-concurrency/general/queue"
)

func main() {
	queue := queue.Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	v, ok := queue.Peek()

	fmt.Println("Front item:", v, ok)
	v, ok = queue.Dequeue()             // Output: 10
	fmt.Println("Dequeue item:", v, ok) // Output: 10
	v, ok = queue.Peek()
	fmt.Println("Front item after dequeue:", v, ok) // Output: 20
	fmt.Println("Is queue empty?", queue.IsEmpty()) // Output: false
}
