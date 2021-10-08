// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	// if strings.Trim(remark, " ") == "" {
	if remark == "" {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) {
		if isYelling(remark) && isWord(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isWord(remark) {
		if isYelling(remark) {
			return "Whoa, chill out!"
		}
	}
	return "Whatever."
}

func isWord(s string) bool {
	return regexp.MustCompile(`[A-Za-z]+`).MatchString(s)
}

func isQuestion(s string) bool {
	return string(s[len(s)-1]) == "?"
}

func isYelling(s string) bool {
	t := strings.ToUpper(s)
	return strings.Compare(s, t) == 0
}
