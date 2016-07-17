package map_function

func MapInt(f func(int) int, input []int) (result [] int) {
	result = make([]int, len(input))
	for i:= range input {
		result[i] = f(input[i])
	}
	return result
}

