package acronym

import (
	"unicode"
)

const testVersion = 2

func Abbreviate(in string) string {
	output := ""
	for i, c := range in {
		if i == 0 {
			output += string(unicode.ToUpper(c))
		} else {
			if (IsCapital(c) && !IsCapital(rune(in[i-1]))) || (unicode.IsLetter(c) && !unicode.IsLetter(rune(in[i-1]))) {
				output += string(unicode.ToUpper(c))
			}
		}
	}
	return output
}

func IsCapital(r rune) bool {
	return unicode.IsLetter(r) && unicode.ToUpper(r) == r
}
