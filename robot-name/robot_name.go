package robotname

import (
	"errors"
	"fmt"
)

var availableNames = map[string]bool{}

type Robot struct {
	name string
}

func (r Robot) hasName() bool {
	return r.name != ""
}

func init() {
	generateNames()
}

func (r *Robot) Name() (string, error) {
	if r.hasName() {
		return r.name, nil
	}

	if len(availableNames) == 0 {
		return "", errors.New("no more available names")
	}

	for k := range availableNames {
		r.name = k
		delete(availableNames, k)
		break
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	name, err := r.Name()
	if err != nil {
		r.name = name
	}
}

func generateNames() {
	var prefix = "AA"
	var number = 0

	for {
		if number > 999 {
			if prefix = increasePrefix(prefix); prefix == "" {
				return
			}

			number = 0
		}

		availableNames[prefix + fmt.Sprintf("%03d", number)] = true
		number++
	}
}

func increasePrefix(prefix string) string {
	runes := []rune(prefix)

	for i := range runes {
		runes[i]++

		if len(runes) - 1 == i && runes[i] > 'Z' {
			return ""
		}

		if runes[i] > 'Z' {
			runes[i] = 'A'
		} else {
			break
		}
	}

	return string(runes)
}