package accumulate

func Accumulate(given []string, converter func(s string) string) []string {
	converted := make([]string, 0, len(given))
	for i := range given {
		converted = append(converted, converter(given[i]))
	}

	return converted
}
