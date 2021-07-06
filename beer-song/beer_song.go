package beer

import (
	"errors"
	"fmt"
)

const upperBoundMax = 99
const lowerBoundMin = 0

func Verse(verse int) (string, error) {
	if verse > upperBoundMax || verse < lowerBoundMin {
		return "", errors.New("incorrect verse")
	}

	var commonParts = "%s of beer on the wall, %s of beer.\n%s, %s of beer on the wall.\n"
	switch true {
	case verse > 2:
		firstTwo := fmt.Sprintf("%d bottles", verse)
		return fmt.Sprintf(commonParts, firstTwo, firstTwo, "Take one down and pass it around", fmt.Sprintf("%d bottles", verse-1)), nil
	case verse == 2:
		return fmt.Sprintf(commonParts, "2 bottles", "2 bottles", "Take one down and pass it around", "1 bottle"), nil
	case verse == 1:
		return fmt.Sprintf(commonParts, "1 bottle", "1 bottle", "Take it down and pass it around", "no more bottles"), nil
	case verse == 0:
		return fmt.Sprintf(commonParts, "No more bottles", "no more bottles", "Go to the store and buy some more", "99 bottles"), nil
	default:
		return "", nil
	}
}

func Verses(upperBound, lowerBound int) (string, error) {
	if upperBound < lowerBound {
		return "", errors.New("incorrect bounds")
	}

	verses := ""
	for i := upperBound; i >= lowerBound; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}

		verses += verse + "\n"
	}

	return verses, nil
}

func Song() string {
	song, err := Verses(upperBoundMax, lowerBoundMin)
	if err != nil {
		return ""
	}

	return song
}
