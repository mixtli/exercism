package accumulate

const testVersion = 1

func Accumulate(input []string, f func(string) string) []string {
	var output []string
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}
