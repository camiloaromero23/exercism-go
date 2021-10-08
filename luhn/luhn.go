package luhn

import (
	"strings"
)

func Valid(input string) bool {
	stripped := strings.Fields(input)
	input = strings.Join(stripped, "")
	input = strings.Join(strings.Split(input, " "), "")
	if len(input) < 2 {
		return false
	}

	aux := []byte(input)
	for _, v := range aux {
		if v < 48 || v > 57 {
			return false
		}
	}
	for i := len(aux) - 2; i >= 0; i -= 2 {
		num := aux[i] - 48
		if num < 0 || num > 9 {
			return false
		}
		num *= 2
		if num > 9 {
			num -= 9
		}
		aux[i] = byte(num + 48)
	}
	sum := 0
	for i := 0; i < len(aux); i++ {
		sum += int(aux[i]) - 48
	}
	return sum%10 == 0
}
