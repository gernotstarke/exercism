package strand

import (
	"strings"
)

var rnamap = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U"}

func ToRNA(dna string) string {
	rna := ""

	for _, s := range strings.Split(dna, "") {
		rna = rna + rnamap[s]
	}

	return rna
}
