package math

func Minimum(values ...int) int {
	min := values[0]
	n := len(values)

	if n == 1 {
		return min
	}

	for i := 1; i < n; i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}

func Maximum(values ...int) int {
	max := values[0]
	n := len(values)

	if n == 1 {
		return max
	}

	for i := 1; i < n; i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}
