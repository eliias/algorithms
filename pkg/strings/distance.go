/*
	Edit distance

	https://iuuk.mff.cuni.cz/~koucky/papers/approxEdit.pdf
	http://www.mit.edu/~andoni/papers/compEdit.pdf
	http://users.monash.edu/~lloyd/tildeAlgDS/Dynamic/Edit/
	https://stackoverflow.com/questions/9452701/ukkonens-suffix-tree-algorithm-in-plain-english
	https://www.cs.helsinki.fi/u/ukkonen/SuffixT1withFigs.pdf
	https://en.wikipedia.org/wiki/Jaccard_index
	https://bitbucket.org/clearer/iosifovich/
	http://www.berghel.net/publications/asm/asm.php
	https://onak.pl/download/publications/Andoni_Krauthgamer_Onak-edit_distance.pdf

	https://en.wikipedia.org/wiki/Vector_space_model

	# String searching
	https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
	https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

	# Hashing & Fingerprint
	https://en.wikipedia.org/wiki/Rabin_fingerprint

	# Local sensitivity hashing

	https://medium.com/engineering-brainly/locality-sensitive-hashing-explained-304eb39291e4
	https://towardsdatascience.com/understanding-locality-sensitive-hashing-49f6d1f6134

	# MinHash
	https://en.wikipedia.org/wiki/MinHash

	# Shingling
	https://nlp.stanford.edu/IR-book/html/htmledition/near-duplicates-and-shingling-1.html
	https://en.wikipedia.org/wiki/W-shingling

	# BK-tree
	https://signal-to-noise.xyz/post/bk-tree/
	https://towardsdatascience.com/symspell-vs-bk-tree-100x-faster-fuzzy-string-search-spell-checking-c4f10d80a078

	# General
	https://en.wikipedia.org/wiki/Metric_space
	Probalistic partial metric spaces
	https://www.researchgate.net/publication/304188311_Fuzzy_Partial_Metric_Spaces
*/
package strings

import (
	"github.com/eliias/algorithms/pkg/math"
	math2 "math"
)

const maxEditDistance = 1

func Hamming(a []rune, b []rune) int {
	if len(a) != len(b) {
		panic("strings are of different lenght")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

func LevenshteinTwoMatrixRows(a []rune, b []rune) int {
	m := len(a)
	n := len(b)

	v0 := make([]int, n+1)
	v1 := make([]int, n+1)

	for i := 0; i <= n; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1

		for j := 0; j < n; j++ {
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			var substitutionCost int
			if a[i] == b[j] {
				substitutionCost = v0[j]
			} else {
				substitutionCost = v0[j] + 1
			}
			v1[j+1] = math.Minimum(deletionCost, insertionCost, substitutionCost)
		}
		v0, v1 = v1, v0
	}

	return v0[n]
}

func Levenshtein(a []rune, b []rune) int {
	if string(a) == string(b) {
		return 0
	}

	if len(a) > len(b) {
		a, b = b, a
	}

	lenA := len(a)
	lenB := len(b)

	for lenA > 0 && a[lenA-1] == b[lenB-1] {
		lenA--
		lenB--
	}

	start := 0
	for start < lenA && a[start] == b[start] {
		start++
	}

	lenA -= start
	lenB -= start

	if lenA == 0 {
		return lenB
	}

	var bRune rune
	var result int
	var temp int
	var temp2 int

	array := make([]int, lenA)
	runes := make([]rune, lenA)
	i := 0
	j := 0

	for i < lenA {
		runes[i] = a[start+i]
		array[i] = i + 1
		i++
	}

	for j < lenB {
		bRune = b[start+j]
		temp = j
		j++
		result = j

		for i = 0; i < lenA; i++ {
			if bRune == runes[i] {
				temp2 = temp
			} else {
				temp2 = temp + 1
			}
			temp = array[i]
			if temp > result {
				if temp2 > result {
					array[i] = result + 1
				} else {
					array[i] = temp2
				}
			} else {
				if temp2 > temp {
					array[i] = temp + 1
				} else {
					array[i] = temp2
				}
			}
			result = array[i]
		}
	}

	return result
}

func LevenshteinIosifovich(subject []rune, query []rune) int {
	// Skip any common prefix.
	subjectBegin, queryBegin := LeftMismatch(subject, 0, len(subject)-1, query, 0, len(query)-1)
	startOffset := subjectBegin - 0

	// If one of the strings is subject prefix of the other, done.
	if len(subject) == startOffset {
		return len(query) - startOffset
	} else if len(query) == startOffset {
		return len(subject) - startOffset
	}

	// Skip any common suffix.
	subjectEnd, _ := LeftMismatch(subject, subjectBegin-1, len(subject)-1, query, queryBegin-1, len(query)-1)
	endOffset := math.Minimum(len(subject)-1-subjectEnd, len(subject)-startOffset)

	// Take the different part.
	subject = subject[startOffset : len(subject)-endOffset-startOffset+1]
	query = query[startOffset : len(query)-endOffset-startOffset+1]

	// Make "subject" the smaller one.
	if len(query) < len(subject) {
		subject, query = query, subject
	}

	// If one of the strings is subject suffix of the other.
	if len(subject) == 0 {
		return len(query)
	}

	// Init buffer.
	buffer := make([]int, len(subject)+1)

	var temp, priorTmp int
	var endJ, startJ, columnMin, p, r int

	for i := 1; i < len(subject)+1; i++ {
		// temp = i - 1
		buffer[0] += 1
		temp = buffer[0]

		startJ = math.Maximum(1, i-maxEditDistance)
		endJ = math.Minimum(len(query)+1, i+maxEditDistance)
		columnMin = maxEditDistance

		for j := startJ; j < endJ; j++ {
			p = temp // p = buffer[j - 1]
			r = buffer[j]

			var cost int
			if subject[i-1] == query[j-1] {
				cost = 0
			} else {
				cost = 1
			}
			temp = math.Minimum(
				math.Minimum(
					r, // Insertion.
					p, // Deletion.
				)+1,
				math.Minimum(
					// Transposition.
					priorTmp+1,
					// Substitution.
					temp+(cost),
				))

			// Keep track of column minimum.
			if temp < columnMin {
				columnMin = temp
			}

			// Record matrix value mat[i -2][j-2]
			priorTmp = temp
			buffer[j], temp = temp, buffer[j]
		}

		if columnMin >= maxEditDistance {
			return maxEditDistance
		}
	}
	return buffer[endJ-1]
}

func EnhancedUkkonen(a []rune, b []rune) int {
	if string(a) == string(b) {
		return 0
	}

	threshold := math2.MaxInt64

	if len(a) > len(b) {
		a, b = b, a
	}

	lenA := len(a)
	lenB := len(b)

	for lenA > 0 && a[lenA-1] == b[lenB-1] {
		lenA--
		lenB--
	}

	if lenA == 0 {
		if lenB < threshold {
			return lenB
		} else {
			return threshold
		}
	}

	tStart := 0
	for tStart < lenA && a[tStart] == b[tStart] {
		tStart++
	}

	lenA -= tStart
	lenB -= tStart

	if lenA == 0 {
		if lenB < threshold {
			return lenB
		} else {
			return threshold
		}
	}

	if lenB < threshold {
		threshold = lenB
	}

	lenD := lenB - lenA

	if threshold < lenD {
		return threshold
	}

	zeroK := int(math2.Floor(math2.Min(float64(threshold), float64(lenA))/2.0)) + 2

	length := lenD + zeroK*2 + 2
	currentRow := make([]int, length+1)
	nextRow := make([]int, length+1)
	for i := 0; i < length; i++ {
		currentRow[i] = -1
		nextRow[i] = -1
	}

	aChars := make([]rune, lenA)
	bChars := make([]rune, lenB)

	i := 0
	t := tStart
	for i < lenA {
		aChars[i] = a[t]
		bChars[i] = b[t]
		i++
		t++
	}

	for i < lenB {
		bChars[i] = b[t]
		i++
		t++
	}

	i = 0
	conditionRow := lenD + zeroK
	endMax := conditionRow * 2

	for {
		i++

		tmp := currentRow
		currentRow = nextRow
		nextRow = tmp

		var start int
		var previousCell int
		currentCell := -1
		var nextCell int

		if i <= zeroK {
			start = -i + 1
			nextCell = i - 2
		} else {
			start = i - zeroK*2 + 1
			nextCell = currentRow[zeroK+start]
		}

		var end int
		if i <= conditionRow {
			end = i
			nextRow[zeroK+i] = -1
		} else {
			end = endMax - 1
		}

		rowIndex := start + zeroK
		k := start
		for k < end {
			previousCell = currentCell
			currentCell = nextCell
			if len(currentRow) == rowIndex+1 {
				break
			}
			nextCell = currentRow[rowIndex+1]

			t := math.Maximum(currentCell+1, previousCell, nextCell+1)

			for t < lenA && t+k < lenB && aChars[t] == bChars[t+k] {
				t++
			}

			nextRow[rowIndex] = t

			k++
			rowIndex++
		}

		if !(nextRow[conditionRow] < lenA && i <= threshold) {
			break
		}
	}

	return i - 1
}

//func DamerauLevenshtein(a []rune, b []rune) int {
//}
//
//func NeedlemanWunsch(a []rune, b []rune) int {
//}
//
//func Hirschberg(a []rune, b []rune) int {
//}
//
//func JaroWinkler(a []rune, b []rune) int {
//}
//
//func Jaccard(a []rune, b []rune) int {
//}
