package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || index > len(slice)-1 {
		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
	if _, ok := GetItem(slice, index); ok {
		slice[index] = value
		return slice
	}
	slice = append(slice, value)
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	slice := make([]int, 0)
	if length > 0 {
		slice = make([]int, length)
	}
	for i := 0; i < len(slice); i++ {
		slice[i] = value
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if _, ok := GetItem(slice, index); !ok {
		return slice
	}
	return append(slice[0:index], slice[index+1:]...)
}
