package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	s := strconv.Itoa(i)
	switch {
	case i == 0:
		return s
	case i % 15 == 0:
		return "FizzBuzz"
	case i % 3 == 0:
		return "Fizz"
	case i % 5 == 0:
		return "Buzz"
	}
	return s
}

