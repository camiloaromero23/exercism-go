package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("not the same lengths")
	}

	n := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
		}
	}

	return n, nil

}
