package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(v int) {
	node := &Node{data: v}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l LinkedList) Display() {
	if l.head == nil {
		fmt.Println("Empty LinkedList")
		return
	}

	current := l.head
	var sb strings.Builder
	for current != nil {
		sb.WriteString(fmt.Sprintf("%d", current.data))
		sb.WriteString("->")
		current = current.next
	}
	sb.WriteString("nil")
	fmt.Println(sb.String())
}

func (l *LinkedList) RemoveDuplicate() {
	if l.head == nil || l.head.next == nil {
		fmt.Println("No duplicate item ")
		return
	}

	m := make(map[int]bool)

	current := l.head
	// cv := current.data
	m[current.data] = true
	pre := current
	current = current.next

	for current != nil {
		_, ok := m[current.data]
		if ok {
			temp := current.next
			// pre.next = current.next
			pre.next = temp
			current = temp

		} else {
			m[current.data] = true
			pre = current
			current = current.next

		}
	}

}

func main() {
	l := LinkedList{}
	l.Add(1)
	l.Add(4)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(3)
	l.Add(4)
	l.Display()
	l.RemoveDuplicate()
	l.Display()
}
