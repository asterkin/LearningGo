package stack

import "testing"

func TestStackEmpty(t *testing.T) {
	s := new(stack)

	if !s.IsEmpty() {
		t.Errorf("Stack Size: expected 0 actual %d", s.Size())
	}

	//TODO Pop should throw Panic
}

func TestStackNonEmpty(t *testing.T) {
	s := new(stack)
	s.Push(125)
	s.Push(257)

	if s.Size() != 2 {
		t.Errorf("Stack Size: expected 2 actual %d", s.Size())
	}

	if s.Top() != 257 {
		t.Errorf("Stack Top: expected 257 actual %d", s.Top())
	}

	v1, v2 := s.Pop(), s.Pop()

	if !s.IsEmpty() {
		t.Errorf("Stack Size: expected 0 actual %d", s.Size())
	}

	if v1 != 257 || v2 != 125 {
		t.Errorf("Stack Pop: expected 257 and 125 actual %d and %d", v1, v2)
	}

}

func TestStackFull(t *testing.T) {
	s := new(stack)
	for i:=0; i<10; i++ {
		s.Push(i)
	}

	if !s.IsFull() {
		t.Errorf("Stack Size: expected %d actual %d", len(s.data), s.Size())

	}

	if s.Top() != 9 {
		t.Errorf("Stack Top: expected 9 actual %d", s.Top())
	}

	//TODO Push should throw Panic

}





