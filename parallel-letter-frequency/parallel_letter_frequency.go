package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	freqCh := make(chan FreqMap)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(s []string, ch chan<- FreqMap, wg *sync.WaitGroup) {
		for _, v := range s {
			ch <- Frequency(v)
		}
		close(ch)
		wg.Done()
	}(s, freqCh, &wg)

	go func(ch <-chan FreqMap, wg *sync.WaitGroup) {
		for freqMap := range ch {
			for k, number := range freqMap {
				if _, ok := m[k]; ok {
					m[k] += number
				} else {
					m[k] = number
				}
			}
		}
		wg.Done()
	}(freqCh, &wg)

	wg.Wait()
	return m
}
