package twelve

import (
	"fmt"
	"strings"
)

var inserts = map[int][2]string{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	total := len(inserts)
	days := make([]string, 0, total)

	for i := 1; i <= total; i++ {
		days = append(days, Verse(i))
	}

	return strings.Join(days, "\n")
}

func Verse(num int) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", inserts[num][0]))

	for i := num; i > 1; i-- {
		sb.WriteString(fmt.Sprintf("%s, ", inserts[i][1]))
	}

	if num > 1 {
		sb.WriteString(fmt.Sprintf("and %s.", inserts[1][1]))
	} else if num == 1 {
		sb.WriteString(fmt.Sprintf("%s.", inserts[1][1]))
	}

	return sb.String()
}
