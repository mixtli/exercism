package diffsquares

import "math"

const testVersion = 1

func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return int(total)
}

func SquareOfSums(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2))
}

func Difference(n int) int {
	return int(SquareOfSums(n) - SumOfSquares(n))
}
