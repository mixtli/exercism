package pythagorean

import "math"

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			c := math.Sqrt(float64(i*i + j*j))
			if c == math.Floor(c) && c <= float64(max) {
				triplets = append(triplets, Triplet{i, j, int(math.Sqrt(float64(i*i + j*j)))})
			}
		}
	}
	return triplets
}

func Sum(p int) []Triplet {
	var triplets []Triplet
	for i := 0; i < p/3; i++ {
		for j := 0; j < (p-i)/2; j++ {
			k := p - i - j
			if i*i+j*j == k*k {
				triplets = append(triplets, Triplet{i, j, k})
			}

		}
	}
	return triplets
}
