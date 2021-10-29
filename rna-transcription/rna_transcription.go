package strand

func ToRNA(dna string) (rna string) {
	rc := map[rune]rune{
		'A': 'U',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}
	for _, n := range dna {
		rna += string(rc[n])
	}
	return rna
}
