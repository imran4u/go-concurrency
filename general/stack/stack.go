package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	size := len(s.items)
	if size == 0 {
		return 0, false
	}

	value := s.items[size-1]
	s.items = s.items[:size-1]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	size := len(s.items)
	if size == 0 {
		return 0, false
	}

	value := s.items[size-1]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
