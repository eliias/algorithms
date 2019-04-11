package strings

import (
	"github.com/eliias/algorithms/pkg/math"
)

// ASCIIControl characters
const ASCIIControl = " "

// ASCIIAlphabetLowercase characters
const ASCIIAlphabetLowercase = "abcdefghijklmnopqrstuvwxyz"

// ASCIIAlphabetUppercase characters
const ASCIIAlphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ASCIINumbers characters
const ASCIINumbers = "0123456789"

// ASCIIPunctuation characters
const ASCIIPunctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// CreateASCIITable maps alphabet to table representation with indices
func CreateASCIITable(alphabet string) map[rune]int {
	m := map[rune]int{}
	for index, r := range []rune(alphabet) {
		m[r] = index
	}
	return m
}

// Vectorize a list of runes by counting character occurrences against given alphabet table
func Vectorize(s []rune, table map[rune]int) math.Vector {
	vector := math.NullVector(len(table))
	for _, r := range s {
		index := table[r]
		vector[index]++
	}
	return vector
}
