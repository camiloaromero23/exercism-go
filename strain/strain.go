package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(pred func(int) bool) (res Ints) {
	for _, v := range i {
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}

func (i Ints) Discard(pred func(int) bool) Ints {
	return i.Keep(
		func(n int) bool {
			return !pred(n)
		})
}

func (s Strings) Keep(pred func(string) bool) (res Strings) {
	for _, v := range s {
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}

func (l Lists) Keep(pred func([]int) bool) (res Lists) {
	for _, v := range l {
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}
