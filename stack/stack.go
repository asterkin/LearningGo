package stack

/*
Package stack implements basic FIFO stack operations
 */
type Stack struct {
	top    int
	data   [10]int
}

//IsEmpty returns true if there are no elements
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

//IsFull returns true if no more elements could be added
func (s *Stack) IsFull() bool {
	return s.top == len(s.data)
}

//Size returns number of elements
func (s *Stack) Size() int {
	return s.top
}

//Top returns the topmost element of the stack
func (s *Stack) Top() int {
	return s.data[s.top - 1]
}

//Push adds a new element to the stack
func (s *Stack) Push(val int) {
	s.data[s.top] = val
	s.top++
}

//Pop removes the topmost element and returns its value
func (s *Stack) Pop() (val int) {
	val = s.Top()
	s.top--
	return
}
