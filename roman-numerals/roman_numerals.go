package romannumerals

import (
	"fmt"
)

var arabic = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

var roman = []string{
	"M",
	"CM",
	"D",
	"CD",
	"C",
	"XC",
	"L",
	"XL",
	"X",
	"IX",
	"V",
	"IV",
	"I",
}

func ToRomanNumeral(n int) (rn string, err error) {
	if n <= 0 || n > 3000 {
		err = fmt.Errorf("number out of range")
		return rn, err
	}
	for i, v := range arabic {
		for n >= v {
			n -= v
			rn += roman[i]
		}
	}
	return rn, err
}
