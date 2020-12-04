package numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{Symbol: "M", Value: 1000},
	{Symbol: "D", Value: 500},
	{Symbol: "C", Value: 100},
	{Symbol: "XC", Value: 90},
	{Symbol: "XCIX", Value: 99},
	{Symbol: "L", Value: 50},
	{Symbol: "XL", Value: 40},
	{Symbol: "X", Value: 10},
	{Symbol: "IX", Value: 9},
	{Symbol: "V", Value: 5},
	{Symbol: "IV", Value: 4},
	{Symbol: "I", Value: 1},
}

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
