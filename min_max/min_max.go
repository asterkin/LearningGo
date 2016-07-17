package min_max

func Min(inp []int) int {
	return find(less, inp)
}

func Max(inp []int) int {
	return find(greater, inp)
}

func find(comp func(int, int) bool, inp []int) (res int) {
	res = inp[0]
	for _,v := range inp[1:] {
		if comp(v, res) {
			res = v
		}
	}
	return
}

func less(v1 int, v2 int) bool {
	return v1 < v2
}

func greater(v1 int, v2 int) bool {
	return v1 > v2
}