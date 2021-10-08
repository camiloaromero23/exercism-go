package reverse

func Reverse(s string) string {
	r := []rune(s)
	var rev []rune
	for i := len(r) - 1; i >= 0; i-- {
		rev = append(rev, r[i])
	}
	return string(rev)
}
