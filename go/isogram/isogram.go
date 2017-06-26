package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(s string) bool {
	found := make(map[rune]bool)
	for _, l := range strings.ToLower(s) {
		if !unicode.IsLetter(rune(l)) {
			continue
		}

		if found[rune(l)] {
			return false
		}
		found[rune(l)] = true
	}
	return true
}
