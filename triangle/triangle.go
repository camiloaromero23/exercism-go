package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func positive(nums ...float64) bool {
	for _, n := range nums {
		if n <= 0 {
			return false
		}
	}
	return true
}

func checkInequality(s1, s2, s3 float64) bool {
	if s1 == math.Inf(+1) || s1 == math.Inf(-1) {
		return false
	}
	return (s1 + s2) >= s3
}

func KindFromSides(a, b, c float64) Kind {
	var ti bool
	cases := [][]float64{{a, b, c}, {b, c, a}, {c, a, b}}
	for i, v := range cases {
		if i == 0 {
			ti = checkInequality(v[0], v[1], v[2])
		} else {
			ti = ti && checkInequality(v[0], v[1], v[2])
		}
	}
	if positive(a, b, c) && ti {
		if a == b && b == c {
			return Equ
		}
		if a == b || a == c || b == c {
			return Iso
		}
		if a != b && b != c && a != c {
			return Sca
		}
	}
	return NaT
}
