package numbers

import (
	"testing"
	"math"
)

func TestFloatAverage(t *testing.T) {
	averageTests := [] struct {
		slice    []float64
		expected float64
	} {
		{[]float64{1.0, 2.5, 3.5}, 2.333},
		{[]float64{0.0, 15.0}, 7.5},
	}

	for _,c := range averageTests {
		actual := FloatAverage(c.slice)
		if math.Abs(actual - c.expected) > 0.001 {
			t.Errorf("FloatAverage(%v expected %v actual %v", c.slice, c.expected, actual)
		}
	}
}

