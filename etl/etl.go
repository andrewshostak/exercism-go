package etl

import "strings"

func Transform(oldMap map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for key, slice := range oldMap {
		for _, v := range slice {
			newMap[strings.ToLower(v)] = key
		}
	}

	return newMap
}
