package sum

func Sum(xs []int) int {
	var result int

	for _, x := range xs {
		result += x
	}

	return result
}

func SumAll(xxs ...[]int) []int {
	// using make
	//result := make([]int, len(xxs))
	//
	//for i, xs := range xxs {
	//  result[i] = Sum(xs)
	//}

	var result []int

	for _, xs := range xxs {
		result = append(result, Sum(xs))
	}

	return result
}

func SumTail(xs []int) int {
	var result int

	if len(xs) < 1 {
		return result
	}

	for _, x := range xs[1:] {
		result += x
	}

	return result
}

func SumTailAll(xxs ...[]int) []int {
	//result := make([]int, len(xxs))
	//
	//for i, xs := range xxs {
	//  result[i] = SumTail(xs)
	//}
	//

	var result []int

	for _, xs := range xxs {
		result = append(result, SumTail(xs))
	}
	return result
}
