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

// SumAllImpure mutates a slice by reassigning indexes on each iteration
// It uses `makr` to create a slice with a capacity upfront
func SumAllImpure(xxs ...[]int) (sums []int) {
	numArgs := len(xxs)
	// slices have a capacity. If we ar assigning values to explicit indexes, we
	// need to know upfront the capacity for the slice. We can then create the
	// slice using `make`
	sums = make([]int, numArgs)

	for i, xs := range xxs {
		sums[i] = Sum(xs)
	}

	return
}

// SumAll will not mutate sums, but will instead reassign the variable on
// each iteration using the `assign` built-in function
func SumAll(xxs ...[]int) (sums []int) {
	sums = []int{}

	for _, xs := range xxs {
		sums = append(sums, Sum(xs))
	}

	return
}
