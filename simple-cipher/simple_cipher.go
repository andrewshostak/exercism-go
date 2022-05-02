package cipher

import (
	"bytes"
	"math"
	"strings"
)

const (
	firstRune = 'a'
	lastRune  = 'z'
)

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if !isValidDistance(distance) {
		return nil
	}

	return shift{distance: int32(distance)}
}

type shift struct {
	distance int32
}

func (c shift) Encode(input string) string {
	lowercased := strings.ToLower(input)
	var b bytes.Buffer
	for _, r := range lowercased {
		if !isValidChar(r) {
			continue
		}

		index := c.distance + r
		if index > lastRune {
			index = firstRune + (index - lastRune - 1)
		} else if index < firstRune {
			index = lastRune - (firstRune - index - 1)
		}
		b.WriteRune(index)
	}

	return b.String()
}

func (c shift) Decode(input string) string {
	c.distance = -(c.distance)
	return c.Encode(input)
}

type vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	var sum int32

	for _, r := range key {
		if !isValidChar(r) {
			return nil
		}
		sum += firstRune - r
	}

	if sum == 0 {
		return nil
	}

	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	return v.vigenere(input, false)
}

func (v vigenere) Decode(input string) string {
	return v.vigenere(input, true)
}

func (v vigenere) vigenere(input string, inverseSign bool) string {
	lowercased := strings.ToLower(input)
	var buffer bytes.Buffer
	var numberOfInvalidChars int
	for i, r := range lowercased {
		if !isValidChar(r) {
			numberOfInvalidChars++
			continue
		}

		distance := getDistance(v.key, i-numberOfInvalidChars)
		if !isValidDistance(distance) {
			buffer.WriteRune(r)
			continue
		}

		if inverseSign {
			distance = -(distance)
		}

		shiftCipher := NewShift(distance)
		if shiftCipher == nil {
			continue
		}

		encoded := shiftCipher.Encode(string(r))
		buffer.WriteString(encoded)
	}

	return buffer.String()
}

func getDistance(key string, i int) int {
	if i > len(key)-1 {
		return getDistance(key, i-len(key))
	}

	return int(key[i]) - firstRune
}

func isValidChar(r rune) bool {
	return r >= firstRune && r <= lastRune
}

func isValidDistance(distance int) bool {
	if distance == 0 {
		return false
	}

	if math.Abs(float64(distance)) > float64(lastRune-firstRune) {
		return false
	}

	return true
}
