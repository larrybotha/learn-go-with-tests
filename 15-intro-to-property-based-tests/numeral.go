package numeral

import "strings"

// MCMLXXXIV => 1984

func ConvertToRoman(n int) string {
	var result strings.Builder

	for n > 0 {
		switch {
		case n > 9:
			result.WriteString("X")
			n -= 10
		case n > 8:
			result.WriteString("IX")
			n -= 9
		case n > 4:
			result.WriteString("V")
			n -= 5
		case n > 3:
			result.WriteString("IV")
			n -= 4
		default:
			result.WriteString("I")
			n--
		}
	}

	return result.String()
}
