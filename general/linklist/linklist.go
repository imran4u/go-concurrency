package linklist

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkList struct {
	head *Node
}

func (l *LinkList) InsertFront(value int) {
	n := Node{value: value}
	if l.head == nil {
		l.head = &n
	} else {
		n.next = l.head
		l.head = &n
	}
}

func (l *LinkList) InsertBack(value int) {
	n := Node{value: value}
	if l.head == nil {
		l.head = &n
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}
	current.next = &n

}

func (l *LinkList) Remove(value int) {

	if l.head == nil {
		fmt.Println("Empty list, Nothing to remove")
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	current := l.head.next
	previous := l.head

	for current != nil && current.value != value {
		previous = current
		current = current.next

	}

	if current == nil {
		fmt.Println("No value found to delete")
	} else {
		previous.next = current.next
	}

}

// VVI
func (l *LinkList) Reverse() {
	if l.head == nil {
		fmt.Println("Empty list, nothing to reverse")
		return
	}
	if l.head.next == nil {
		fmt.Println("only one element, nothing to reverse")
		return
	}
	//
	var prev *Node
	current := l.head
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}

	l.head = prev
}

func (l *LinkList) Display() {
	if l.head == nil {
		fmt.Println("Empty link list, nothing to print")
		return
	}

	var sb strings.Builder
	current := l.head

	for current != nil {
		sb.WriteString(fmt.Sprintf("%d", current.value))
		sb.WriteString("->")
		current = current.next
	}

	fmt.Println(strings.TrimSuffix(sb.String(), "->"))

}
