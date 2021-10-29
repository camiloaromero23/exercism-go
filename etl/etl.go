package etl

import "strings"

func Transform(legacy map[int][]string) map[string]int {
	shinyScore := make(map[string]int)
	for s, letters := range legacy {
		for _, l := range letters {
			l = strings.ToLower(l)
			shinyScore[l] = s
		}
	}
	return shinyScore
}
