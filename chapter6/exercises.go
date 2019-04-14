package exercises

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

// Oh, duh. max.
func greatest(xs ...int) int {
	var greatest int
	for i, e := range xs {
		if i == 0 || e > greatest {
			greatest = e
		}
	}
	return greatest
}

func makeOddGenerator() func() int {
	x := -1
	return func() int {
		x = x + 2
		return x
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// ohh right, multiple variable assignments. nice.
func swap(x *int, y *int) {
	// Book solution:
	// *x, *y = *y, *x
	temp := *x
	*x = *y
	*y = temp
}
