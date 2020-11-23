package protein

import "errors"

var (
	ErrStop        = errors.New("codon stop error")
	ErrInvalidBase = errors.New("invalid base error")
)

func FromRNA(input string) (proteins []string, err error) {
	for i := 0; i < len(input); i += 3 {
		cobon := input[i : i+3]
		protein, err := FromCodon(cobon)
		if err == ErrStop {
			return proteins, nil
		}
		if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, err
}

func FromCodon(input string) (protein string, err error) {
	switch input {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}
	return protein, err
}
