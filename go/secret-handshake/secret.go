package secret

import "math"

const testVersion = 1

// Handshake decodes an integer into an array of actions
func Handshake(n uint) []string {
	var output []string
	code := map[uint]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}
	for i := 0; i <= 3; i++ {
		place := uint(math.Pow(2, float64(i)))
		if (place & n) > 0 {
			output = append(output, code[place])
		}
	}
	if (16 & n) > 0 {
		reverse(output)
	}
	return output
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
