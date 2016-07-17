package sort

func Sort(inp []int) []int {
	l := len(inp)
	if l < 2 {
		return inp
	}
	for i:= range inp[:l-1] {
		for j := i+1; j<l; j++ {
			if inp[i] > inp[j] {
				inp[i], inp[j] = inp[j], inp[i]
			}
		}
	}
	return inp
}

