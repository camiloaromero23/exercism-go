package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap)
	defer close(c)

	for _, w := range s {
		go func(w string) {
			c <- Frequency(w)
		}(w)
	}

	out := <-c

	for i := 1; i < len(s); i++ {
		for k, v := range <-c {
			out[k] += v
		}
	}
	return out

}
