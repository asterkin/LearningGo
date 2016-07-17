package main

import (
	"LearningGo/greeting"
	"fmt"
	"os"
	"LearningGo/fizzbuzz"
	"LearningGo/fibonacci"
	"LearningGo/map_function"
	"LearningGo/min_max"
)

func main() {
	fmt.Println(greeting.Greeting(os.Args[1:]))
	for i:=0; i<100; i++ {
		fmt.Print(fizzbuzz.FizzBuzz(i))
		if (i + 1) % 25 == 0 {
			fmt.Println()
		} else {
			fmt.Print("\t")
		}
	}
	fmt.Println(fibonacci.Fibonacci(20))
	fmt.Println(map_function.MapInt(triple, []int{1,2,3}))
	input := [] int {11,75, 28, 36, 19}
	fmt.Println(min_max.Min(input), min_max.Max(input))
}

func triple(i int) int {
	return i * 3
}
