package iteration

import "strings"

func RepeatNaive(x string, n int) string {
	var result string

	for range n {
		result += x
	}

	return result
}

func Repeat(x string, n int) string {
	// strings.Builder is more performant when there are many
	// concatenations
	var result strings.Builder

	for range n {
		result.WriteString(x)
	}

	return result.String()
}
