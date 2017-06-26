package hamming

import (
	"errors"
	"fmt"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strings must be the same length")
	}
	l := len(a)
	d := 0
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
