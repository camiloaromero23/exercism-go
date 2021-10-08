// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about
// h. t. s://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = fixSpecialWords(s)
	arrS := strings.Fields(s)

	var newS string
	for _, c := range arrS {
		newS = newS + strings.ToUpper(string(c[0]))
	}
	return newS
}

func fixSpecialWords(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	return s
}
