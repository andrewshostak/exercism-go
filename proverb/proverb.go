package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var output []string
	for a, b := 0, 1; b < len(rhyme); {
		output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[a], rhyme[b]))
		a++
		b++
	}

	if len(rhyme) > 0 {
		output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return output
}
