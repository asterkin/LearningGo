package numbers

func FloatSum(vals []float64) (sum float64) {
	sum = 0
	for _,v := range vals {
		sum += v
	}
	return
}

func FloatAverage(vals []float64) float64 {
	return FloatSum(vals) / float64(len(vals))
}

