package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(parts []string) FreqMap {
	frequencyChannel := make(chan FreqMap, len(parts))

	for _, part := range parts {
		go func(part string, ch chan<- FreqMap) {
			ch <- Frequency(part)
		}(part, frequencyChannel)
	}

	frequencyMap := FreqMap{}
	for range parts {
		for k, v := range <-frequencyChannel {
			frequencyMap[k] += v
		}
	}

	return frequencyMap
}
