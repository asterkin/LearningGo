package rpn

import (
	"testing"
	"bufio"
	"strings"
)

func TestRpn(t *testing.T) {
	rpnTests := []struct {
		input    string
		expected int
	} {
		{"",              0},
		{"123\n",       123},
		{"123 255 +\n", 378},
		{"123 255 -\n",-132},
		{"123 10 *\n", 1230},
		{"123 10 /\n",   12},
		{"123\n255\n+\n", 378},
	}
	for _,c := range rpnTests {
		reader := bufio.NewReader(strings.NewReader(c.input))
		calc   := NewCalc()
		for s,err := reader.ReadString('\n'); err == nil; s,err = reader.ReadString('\n') {
			calc.Process(s)
		}
		actual := calc.Result()
		if actual != c.expected {
			t.Errorf("RPN Calculator(%s, %v) expected: %d actual: %d", c.input, calc, c.expected, actual)
		}
	}
}

