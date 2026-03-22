package task03collections

func WordFrequency(words []string) map[string]int {
	freq := make(map[string]int, len(words))
	for _, word := range words {
		freq[word]++
	}

	return freq
}
