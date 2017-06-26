package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	letters := make(map[rune]bool)
	for _, l := range strings.ToLower(s) {
		if l >= 'a' && l <= 'z' {
			letters[l] = true
		}
	}
	return len(letters) == 26
}
