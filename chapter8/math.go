package math

import "errors"

// Min - return the smallest number in the slice
func Min(xs []float64) (float64, error) {
	if xs == nil || len(xs) == 0 {
		return 0, errors.New("Empty slice")
	}

	var min float64
	for i, e := range xs {
		if i == 0 || e < min {
			min = e
		}
	}
	return min, nil
}

// Max - return the greatest number in the slice
func Max(xs []float64) (float64, error) {
	if xs == nil || len(xs) == 0 {
		return 0, errors.New("Empty slice")
	}

	var max float64
	for i, e := range xs {
		if i == 0 || e > max {
			max = e
		}
	}
	return max, nil
}
