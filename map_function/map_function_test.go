package map_function

import (
	"testing"
	"fmt"
)

func TestMapFunction(t *testing.T) {
	input := make([] int, 10)
	for i := range input {
		input[i] = i
	}
	mapFunctionTests := [] struct {
		function func(int) int
		expected [] int
	} {
		{add(1), []int{1,2,3,4,5,6,7,8,9,10}},
		{mult(2),[]int{0,2,4,6,8,10,12,14,16,18}},
	}

	for _,c := range mapFunctionTests {
		actual := MapInt(c.function, input)
		if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", c.expected) {
			t.Errorf("MapInt(%v) expected: %v actual: %v ", input, c.expected, actual)
		}
	}
}

func add(v1 int) func(int) int {
	return func(v2 int) int {
		return v1 + v2
	}
}

func mult(v1 int) func(int) int {
	return func(v2 int) int {
		return v1 * v2
	}
}



