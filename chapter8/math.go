package math

// Min - return the smallest number in the slice
func Min(xs []float64) float64 {
	var min float64
	for i, e := range xs {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}

// Max - return the greatest number in the slice
func Max(xs []float64) float64 {
	var max float64
	for i, e := range xs {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}
