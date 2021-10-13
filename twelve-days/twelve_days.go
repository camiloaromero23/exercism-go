package twelve

import "fmt"

var (
	opening = "On the %s day of Christmas my true love gave to me: "
	verses  = map[int]string{
		1:  "and a Partridge in a Pear Tree.",
		2:  "two Turtle Doves",
		3:  "three French Hens",
		4:  "four Calling Birds",
		5:  "five Gold Rings",
		6:  "six Geese-a-Laying",
		7:  "seven Swans-a-Swimming",
		8:  "eight Maids-a-Milking",
		9:  "nine Ladies Dancing",
		10: "ten Lords-a-Leaping",
		11: "eleven Pipers Piping",
		12: "twelve Drummers Drumming",
	}
	numbering = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
)

func Song() (s string) {
	for i := range numbering {
		s += Verse(i + 1)
		if i < 11 {
			s += "\n"
		}
	}
	return
}

func Verse(input int) string {
	v := fmt.Sprintf(opening, numbering[input-1])
	if input == 1 {
		return v + "a Partridge in a Pear Tree."
	}
	for i := input; i > 0; i-- {
		v += verses[i]
		if i > 1 {
			v += ", "
		}
	}
	return v
}
