package bob

import (
	"unicode"
)

const testVersion = 3

func Hey(s string) string {
	is_question := false
	is_shout := true
	silence := true
	has_letters := false
	for _, c := range s {
		if !IsWhitespace(c) {
			silence = false
			is_question = false
		}

		if c == '?' {
			is_question = true
		}

		if unicode.IsLetter(c) {
			has_letters = true
		}
		if unicode.IsLetter(c) && !IsCapital(c) {
			is_shout = false
		}
	}
	if silence {
		return "Fine. Be that way!"
	}

	if is_shout && has_letters {
		return "Whoa, chill out!"
	}

	if is_question {
		return "Sure."
	}
	return "Whatever."
}

func IsCapital(r rune) bool {
	return unicode.IsLetter(r) && unicode.ToUpper(r) == r
}

func IsWhitespace(r rune) bool {
	if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
		return true
	} else {
		return false
	}
}
