package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	out := []string{}
	for i := 0; i < len(s) && (i+n <= len(s)); i++ {
		out = append(out, s[i:i+n])
	}

	return out
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return UnsafeFirst(n, s), true
}
