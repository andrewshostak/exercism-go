package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences should be equal length")
	}

	if len(a) == 0 {
		return 0, nil
	}

	total := 0
	for i := range a {
		if a[i] != b[i] {
			total++
		}
	}

	return total, nil
}
