package protein

import "errors"

var ErrStop = errors.New("stop codon received")
var ErrInvalidBase = errors.New("invalid codon received")

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	var lastCodon string
	for _, v := range rna {
		lastCodon += string(v)

		if len(lastCodon) == 3 {
			protein, err := FromCodon(lastCodon)

			if err == ErrStop {
				return proteins, nil
			}

			if err == ErrInvalidBase {
				return []string{}, err
			}

			proteins = append(proteins, protein)
			lastCodon = ""
		}
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
