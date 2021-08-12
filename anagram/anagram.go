package anagram

import "strings"

func Detect(input string, candidates []string) []string {
	var matched []string
	prepared := prepareString(input)
	inputMap := stringToMap(prepared)

	for i := range candidates {
		candidatePrepared := prepareString(candidates[i])

		if prepared == candidatePrepared {
			continue
		}

		candidateMap := stringToMap(candidatePrepared)

		if compareMap(inputMap, candidateMap) {
			matched = append(matched, candidates[i])
		}
	}

	return matched
}

func compareMap(first, second map[byte]int) bool {
	if len(first) != len(second) {
		return false
	}

	for key := range first {
		if _, ok := second[key]; !ok {
			return false
		}

		if first[key] != second[key] {
			return false
		}
	}

	return true
}

func stringToMap(input string) map[byte]int {
	m := make(map[byte]int)

	for i := range input {
		m[input[i]]++
	}

	return m
}

func prepareString(s string) string {
	return strings.ToLower(s)
}
