package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	fizzBuzzTests := []struct {
		i        int
		expected string
	} {
		{0, "0"},
		{3, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{15, "FizzBuzz"},
	}

	for _, c := range fizzBuzzTests {
		actual := FizzBuzz(c.i)
		if actual != c.expected {
			t.Errorf("FizBuzz(%d) expected '%s' actual '%s'", c.i, c.expected, actual)
		}
	}
}
