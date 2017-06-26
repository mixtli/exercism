package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}
	frequencies := make(chan FreqMap, len(input))
	for _, s := range input {
		go func(s string) {
			frequencies <- Frequency(s)
		}(s)
	}
	for range input {
		for letter, count := range <-frequencies {
			m[letter] += count
		}
	}
	return m
}
