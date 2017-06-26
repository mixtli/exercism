package grains

import (
	"errors"
	"math"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Out of Range")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

func Total() uint64 {
	var total uint64
	var val uint64
	for i := 1; i <= 64; i++ {
		val, _ = Square(i)
		total += val
	}
	return total
}
