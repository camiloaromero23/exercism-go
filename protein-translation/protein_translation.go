package protein

import "errors"

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("STOP")

var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromCodon(codon string) (protein string, err error) {
	protein, ok := proteins[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return
}

func FromRNA(rna string) (proteinSlice []string, err error) {
	if len(rna) < 3 {
		err = errors.New("Invalid rna strand")
		return
	}
	proteinSlice = make([]string, 0, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return proteinSlice, err
		}
		proteinSlice = append(proteinSlice, protein)
	}
	return proteinSlice, nil
}
