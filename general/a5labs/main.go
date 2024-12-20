package main

import (
	"fmt"
	"strings"
)

// Problem Title: Intersection of Two Linked Lists
// Description:
// Given the heads of two singly linked lists, determine if the two lists intersect at some node.
// The intersection means that the two linked lists converge at some point and continue as a single
//  sequence thereafter. Return the true if there are intersected. If the two linked lists do not intersect,
//  return false

// example1
// L1 :3 - 4 -5 -10 - 12 - 15  true
// 	       /
// L2  1 -2

// example2
// 2-5-7
// 3-10-2-12  false

type Node struct {
	Value int
	next  *Node
}

type LinkList struct {
	head *Node
}

func (l *LinkList) InsertNode(node *Node) {
	if l.head == nil {
		l.head = node
		return
	}

	// add at end
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkList) DisPlay() {
	if l.head == nil {
		fmt.Println("Nothing to display")
		return
	}

	var sb strings.Builder
	current := l.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%d", current.Value))
		sb.WriteString("->")
		current = current.next
	}
	sb.WriteString("nil")
	fmt.Println(sb.String())

}

func insersect(l1, l2 LinkList) bool {
	if l1.head == nil || l2.head == nil {
		return false
	}
	c1 := l1.head

	// #first sol : time O(n2)
	// for c1 != nil {
	// 	c2 := l2.head
	// 	for c2 != nil {
	// 		if c1 == c2 {
	// 			return true
	// 		}
	// 		c2 = c2.next
	// 	}
	// 	c1 = c1.next
	// }

	// #second sol : time O(n) but memory O(n)
	// m := make(map[*Node]bool)
	// for c1 != nil {

	// 	c1 = c1.next
	// 	m[c1] = true
	// }

	// c2 := l2.head
	// for c2 != nil {
	// 	_, ok := m[c2]
	// 	if ok {
	// 		return true
	// 	}
	// 	c2 = c2.next

	// }
	// return false

	// # third sol: time O(n)
	for c1.next != nil {
		c1 = c1.next
	}
	c2 := l2.head
	for c2.next != nil {
		c2 = c2.next
	}

	return c1 == c2

}

func main() {

	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node10 := &Node{Value: 10}
	//3 -> 4 ->5 ->10
	l1 := LinkList{}
	l1.InsertNode(node3)
	l1.InsertNode(node4)
	l1.InsertNode(node5)
	l1.InsertNode(node10)
	l1.DisPlay()

	// 1 -2-4-5-10
	l2 := LinkList{}
	l2.InsertNode(node1)
	l2.InsertNode(node2)
	l2.InsertNode(node4) // intersection

	l2.DisPlay()

	fmt.Println("Is intersect = ", insersect(l1, l2))

}
