package main

import (
	"fmt"
	"strings"
)

func main() {
	l := LinkList{}
	l.display() //
	l.insert(12)
	l.insert(15)
	l.insert(25)
	l.display() //>>>

	// l.remove(15)
	// l.display()

}

type Node struct {
	value int
	next  *Node
}

type LinkList struct {
	head *Node
}

func (l *LinkList) insert(v int) {
	node := Node{value: v}

	if l.head == nil {
		l.head = &node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node
	}

}

func (l *LinkList) remove(value int) {

	if l.head == nil {
		fmt.Println("Empty linked list nothing to remvoe")
	} else {
		//TODO:
	}

}
func (l *LinkList) display() {

	if l.head == nil {
		fmt.Println("Nothing to print empty linked list")
	} else {
		var builder strings.Builder
		current := l.head

		for current != nil {
			builder.WriteString(fmt.Sprintf("%d", current.value))
			builder.WriteString(", ")
			current = current.next
		}

		fmt.Println(builder.String()[:len(builder.String())-2])
	}

}
