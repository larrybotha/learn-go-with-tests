package sum

// Sum5 expects an array of 5 integers
func Sum5(xs [5]int) int {
	sum := 0

	//for i := 0; i < len(xs); i++ {
	//sum += xs[i]
	//}

	for _, number := range xs {
		sum += number
	}

	return sum
}

// Sum expects an int slice
func Sum(xs []int) int {
	sum := 0

	for _, number := range xs {
		sum += number
	}

	return sum
}
