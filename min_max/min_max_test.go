package min_max

import "testing"

var input = []int{10,2,7,5,1,79,0,126}

func TestMin(t *testing.T) {
	Min([]int{5})
	actual := Min(input)
	if actual != 0 {
		t.Errorf("Min(%v) expected: 0 actual: %d", input, actual)
	}
}

func TestMax(t *testing.T) {
	Max([]int{5})
	actual := Max(input)
	if actual != 126 {
		t.Errorf("Max(%v) expected: 0 actual: %d", input, actual)
	}
}
