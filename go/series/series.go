package series

const testVersion = 2

func All(n int, s string) []string {
	var strs []string
	for i := 0; i < len(s)-n+1; i++ {
		strs = append(strs, s[i:i+n])
	}
	return strs
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	ok = n <= len(s)
	if ok {
		first = UnsafeFirst(n, s)
	}
	return first, ok
}
