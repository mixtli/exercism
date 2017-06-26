package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

func CanQueenAttack(w string, b string) (bool, error) {
	if w == b {
		return false, errors.New("Can not be same square")
	}
	w1 := w[0] - 97
	w2 := w[1] - 49
	b1 := b[0] - 97
	b2 := b[1] - 49
	for _, i := range []byte{w1, w2, b1, b2} {
		if i < 0 || i > 7 {
			return false, errors.New("Out of Range")
		}
	}
	if math.Abs(float64(int(w1)-int(b1))) == math.Abs(float64(int(w2)-int(b2))) {
		return true, nil
	}
	if w1 == b1 || w2 == b2 {
		return true, nil
	}

	return false, nil
}
