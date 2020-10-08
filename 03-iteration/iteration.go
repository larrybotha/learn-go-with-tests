package iteration

func Repeat(value string, repeats int) string {
	var result string

	for n := 0; n < repeats; n++ {
		result = result + value
	}

	return result
}
