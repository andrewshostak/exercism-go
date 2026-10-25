package parallelletterfrequency

import (
	"strings"
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	s = strings.ToLower(s)
	frequencyMap := FreqMap{}

	runes := []rune(s)
	for _, r := range runes {
		if r < 'a' {
			continue
		}

		frequencyMap[r]++
	}

	return frequencyMap
}

func ConcurrentFrequency(parts []string) FreqMap {
	total := FreqMap{}

	ch := make(chan FreqMap, len(parts))
	for _, part := range parts {
		go func(part string) {
			ch <- Frequency(part)
		}(part)
	}

	for range parts {
		for k, v := range <-ch {
			total[k] += v
		}
	}

	return total
}
