package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	startLower, endLower := 'a', 'z'
	startUpper, endUpper := 'A', 'Z'
	output := []rune(plain)
	for i, r := range plain {
		// lowercase
		if r >= startLower && r <= endLower {
			char := rotate(r, rune(shiftKey), startLower, endLower)
			output[i] = char
			continue
		}

		// uppercase
		if r >= startUpper && r <= endUpper {
			char := rotate(r, rune(shiftKey), startUpper, endUpper)
			output[i] = char
			continue
		}

		output[i] = r
	}

	return string(output)
}

func rotate(letter, shiftKey, start, end rune) rune {
	fullCycle := end - start + 1
	return (letter+shiftKey-start)%fullCycle + start
}
