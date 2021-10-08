package isogram

import (
	"strings"
)

func IsIsogram(w string) bool {

	w = strings.ToLower(w)

	for _, v := range w {
		if strings.Count(w, string(v)) > 1 && string(v) != " " && string(v) != "-" {
			return false
		}
	}

	return true
}
