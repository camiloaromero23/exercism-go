package grains

import (
	"fmt"
	"math"
)

func Square(n int) (uint64, error) {

	if n <= 0 {
		return 0, fmt.Errorf("Error: Cannot be 0 or less")
	}
	if n > 64 {
		return 0, fmt.Errorf("Error: Cannot be greater than 64")
	}
	return uint64(math.Pow(2, float64(n)-1)), nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		sum += val
	}
	return sum
}
