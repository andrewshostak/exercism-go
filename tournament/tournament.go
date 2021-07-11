package tournament

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	. "tournament/rtng"
)

func Tally(r io.Reader, w io.Writer) error {
	var rating []RatingItem

	buf, err := ioutil.ReadAll(r)
	lines := bytes.Split(buf, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		lineParts := bytes.Split(line, []byte(";"))
		if len(lineParts) != 3 {
			return errors.New("invalid input")
		}

		result := string(lineParts[2])
		if err = validateResult(result); err != nil {
			return err
		}

		rating = UpdateRating(rating, string(lineParts[0]), string(lineParts[1]), result)
	}

	SortRating(rating)
	_, err = w.Write([]byte(StringifyTable(rating)))

	return err
}

func validateResult(result string) error {
	if result != "win" && result != "draw" && result != "loss" {
		return errors.New("invalid result")
	}

	return nil
}
