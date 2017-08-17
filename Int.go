package gohelpmath

// IntMin returns the lowest value of the provided integers
func IntMin(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

// IntMax returns the highest value of the provided integers
func IntMax(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// Abs returns the absolute value (Always positive) of the integer
func Abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

// Abs64 returns the absolute value (Always positive) of the integer
func Abs64(value int64) int64 {
	if value < 0 {
		return value * -1
	}
	return value
}
