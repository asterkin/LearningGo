package stack

type stack struct {
	top    int
	data   [10]int
}

func (s *stack) IsEmpty() bool {
	return s.top == 0
}

func (s *stack) IsFull() bool {
	return s.top == len(s.data)
}

func (s *stack) Size() int {
	return s.top
}

func (s *stack) Top() int {
	return s.data[s.top - 1]
}

func (s *stack) Push(val int) {
	s.data[s.top] = val
	s.top++
}

func (s *stack) Pop() (val int) {
	val = s.Top()
	s.top--
	return
}
