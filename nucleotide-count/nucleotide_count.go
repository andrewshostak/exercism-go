package dna

import "errors"

// Histogram is a mapping from Nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[Nucleotide]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

type Nucleotide rune

const (
	Adenine  Nucleotide = 'A'
	Cytosine Nucleotide = 'C'
	Guanine  Nucleotide = 'G'
	Thymine  Nucleotide = 'T'
)

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid Nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{Adenine: 0, Cytosine: 0, Guanine: 0, Thymine: 0}

	for _, r := range d {
		nucleotide := Nucleotide(r)
		if _, ok := h[nucleotide]; !ok {
			return nil, errors.New("invalid Nucleotide type")
		}

		h[nucleotide]++
	}

	return h, nil
}
