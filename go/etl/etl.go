package etl

import "strings"

const testVersion = 1

// Transform converts a structure like {1: {"A", "B"}} to {"a": 1, "b": 1}
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		for i := 0; i < len(v); i++ {
			output[strings.ToLower(v[i])] = k
		}

	}
	return output
}
