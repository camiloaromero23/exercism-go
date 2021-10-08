package raindrops

import (
	"fmt"
)

func Convert(n int) string {

	s := ""
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return fmt.Sprint(n)
	}

	return s
}
