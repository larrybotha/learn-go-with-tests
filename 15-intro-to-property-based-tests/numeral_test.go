package numeral

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("converts numbers to roman equivalents", func(t *testing.T) {
		cases := []struct {
			Name  string
			Input int
			Want  string
		}{
			{"converts 1 to I", 1, "I"},
			{"converts 2 to II", 2, "II"},
			{"converts 3 to III", 3, "III"},
			{"converts 4 to IV", 4, "IV"},
			{"converts 5 to V", 5, "V"},
			{"converts 6 to VI", 6, "VI"},
			{"converts 7 to VII", 7, "VII"},
			{"converts 8 to VIII", 8, "VIII"},
			{"converts 9 to IX", 9, "IX"},
			{"converts 10 to X", 10, "X"},

			{"14 gets converted to XIV", 14, "XIV"},
			{"18 gets converted to XVIII", 18, "XVIII"},
			{"20 gets converted to XX", 20, "XX"},
			{"39 gets converted to XXXIX", 39, "XXXIX"},
		}

		for _, v := range cases {
			t.Run(v.Name, func(t *testing.T) {
				got := ConvertToRoman(v.Input)
				want := v.Want

				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			})
		}
	})
}
