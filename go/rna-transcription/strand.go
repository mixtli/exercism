package strand

const testVersion = 3

var xlate = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA translates a strand of RNA to DNA
func ToRNA(dna string) string {
	rna := ""
	for _, n := range dna {
		rna += string(xlate[n])
	}
	return rna
}
