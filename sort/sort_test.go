package sort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	sortTests := []struct {
		input    []int
		expected []int
	}{
		{[]int{},                  []int{}},
		{[]int{128},               []int{128}},
		{[]int{128,1,127,2,126,3}, []int{1,2,3,126,127,128}},
	}

	for _,c := range sortTests {
		actual := Sort(c.input)
		if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", c.expected) {
			t.Errorf("Sort expected: %v actual: %v", c.expected, actual)
		}
	}
}

