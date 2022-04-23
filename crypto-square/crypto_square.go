package cryptosquare

import (
	"bytes"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// lowercase
	lowercased := strings.ToLower(pt)

	// keep only letters and numbers
	var b bytes.Buffer
	for _, r := range lowercased {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		}
	}
	input := b.String()

	// get number of columns and rows
	cols, rows := getColsAndRows(len(input))

	// add spaces if necessary
	if numOfSpaces := (cols * rows) - len(input); numOfSpaces > 0 && len(input) > 0 {
		spaces := strings.Repeat(" ", numOfSpaces)
		input += spaces
	}

	// encode
	encoded := make([]byte, 0, len(input)+cols-1)
	i, offset := 0, 0
	for offset < cols {
		index := i*cols + offset
		if index >= len(input) {
			offset++
			i = 0
			if offset < cols {
				encoded = append(encoded, ' ')
			}
		} else {
			encoded = append(encoded, input[index])
			i++
		}
	}

	return string(encoded)
}

func getColsAndRows(length int) (int, int) {
	cols, rows := 1, 1
	for (cols * rows) < length {
		if cols == rows {
			cols++
		} else {
			rows++
		}
	}

	return cols, rows
}
