package protein

const testVersion = 1

var codons = map[string]string{
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

// FromCodon returns an amino acid given a three nucleotide codon
func FromCodon(codon string) string {
	return codons[codon]
}

// FromRNA returns an array of amino acids given a strand of RNA
func FromRNA(strand string) []string {
	var proteins []string
	for i := 0; i < len(strand); i += 3 {
		codon := strand[i : i+3]
		if FromCodon(codon) == "STOP" {
			break
		}
		proteins = append(proteins, FromCodon(codon))
	}
	return proteins
}
