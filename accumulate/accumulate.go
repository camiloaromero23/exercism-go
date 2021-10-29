package accumulate

func Accumulate(in []string, converter func(string) string) (acc []string) {
	for _, v := range in {
		acc = append(acc, converter(v))
	}
	return acc
}
