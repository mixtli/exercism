package triangle

import "math"

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind uint8

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	if !(a > 0 && b > 0 && c > 0) {
		return NaT
	}

	if a+b < c || b+c < a || c+a < b {
		return NaT
	}
	for _, l := range []float64{a, b, c} {
		if l == math.Inf(1) || l == math.Inf(-1) || l == math.NaN() {
			return NaT
		}
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
