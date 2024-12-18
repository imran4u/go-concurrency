package main

import "github.com/imran4u/go-concurrency/general/linklist"

func main() {
	l := linklist.LinkList{}
	l.InsertFront(5)
	l.InsertFront(4)
	l.InsertFront(3)
	l.Display()

	l.InsertBack(6)
	l.InsertBack(7)
	l.InsertBack(8)
	l.Display()

	// l.Remove(7)
	// l.Remove(8)
	// l.Remove(3)
	l.Remove(14)
	l.Display()

}
