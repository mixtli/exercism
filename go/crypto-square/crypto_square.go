package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode a message using a cryptosquare
func Encode(message string) string {
	message = normalize(message)
	numRows := int(math.Ceil(math.Sqrt(float64(len(message)))))
	rows := make([]string, numRows)
	for i, char := range message {
		rows[i%numRows] += string(char)
	}
	return strings.Join(rows, " ")

}

func normalize(message string) string {
	newStr := ""
	for _, c := range message {
		if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) {
			newStr += string(unicode.ToLower(rune(c)))
		}
	}
	return newStr
}
