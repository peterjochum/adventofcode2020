package utility

// Min returns the lowest num in numbers
func Min(numbers []int) int {
	var min int
	for i, n := range numbers {
		if i == 0 || n < min {
			min = n
		}
	}
	return min
}

// Max returns the largest num in numbers
func Max(numbers []int) int {
	var max int
	for i, n := range numbers {
		if i == 0 || n > max {
			max = n
		}
	}
	return max
}
