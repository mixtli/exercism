package summultiples

const testVersion = 2

// SumMultiples Finds the sum of all multiples of given divisors
func SumMultiples(limit int, divisors ...int) int {
	multiples := []int{}
	total := 0
	for i := 0; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				total += i
				multiples = append(multiples, i)
				break
			}
		}
	}
	return total
}
