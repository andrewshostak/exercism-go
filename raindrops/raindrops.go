package raindrops

import (
	"sort"
	"strconv"
)

func Convert(input int) (raindrops string) {
	rainDropsMap := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	keys := make([]int, 0, len(rainDropsMap))

	for key := range rainDropsMap {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		if input % key == 0 {
			raindrops += rainDropsMap[key]
		}
	}

	if raindrops == "" {
		raindrops = strconv.Itoa(input)
	}

	return
}