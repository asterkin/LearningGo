package fibonacci

func Fibonacci(count int) (result []int) {
	result = make([]int, count)
	if count < 2 {
		return
	}
	result[1] = 1
	for i:=2; i<count; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return
}

