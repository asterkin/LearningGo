package fibonacci

import (
	"testing"
	"fmt"
)

func TestFibonacci(t *testing.T) {
	fibonacciTests := [] struct {
		count    int
		expected []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{2, []int{0,1}},
		{8, []int{0,1,1,2,3,5,8,13}},
	}

	for _,c := range fibonacciTests {
		actual := Fibonacci(c.count)
		if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", c.expected) {
			t.Errorf("Fibonacci(%d) expected %v actual %v", c.count, c.expected, actual)
		}
	}
}

