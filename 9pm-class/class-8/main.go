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
	l.insert(35)
	l.insert(45)

	l.display() //>>>

	l.remove(45)
	l.display()

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

func (l *LinkList) remove(v int) {

	if l.head == nil {
		fmt.Println("Empty linked list nothing to remvoe")
		return
	}

	if l.head.value == v {
		fmt.Println("deleting head value =", l.head.value)
		l.head = l.head.next
		return
	}

	current := l.head

	for current.next != nil && current.next.value != v {
		current = current.next
	}

	if current.next.value == v {
		fmt.Println("deleting value =", current.next.value)
		current.next = current.next.next
	} else {
		fmt.Println("Value not found")
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
			builder.WriteString("-> ")
			current = current.next
		}
		dispaly := builder.String()[:len(builder.String())-3]
		fmt.Println(dispaly + "->Nil")
	}

}
