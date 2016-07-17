package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	s := fmt.Sprintf("%d", i)
	if i == 0 {
		return s
	}
	if i % 15 == 0 {
		return "FizzBuzz"
	}
	if i % 3 == 0 {
		return "Fizz"
	}
	if i % 5 == 0 {
		return "Buzz"
	}
	return s
}

