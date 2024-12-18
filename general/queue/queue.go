package queue

type Queue struct {
	items []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // Queue is empty
	}
	frontItem := q.items[0]
	q.items = q.items[1:] // Remove the first element
	return frontItem, true
}

// Peek returns the front element without removing it
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
