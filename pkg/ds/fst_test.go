package ds

import (
	"fmt"
	"testing"
)

func TestFST_Add(t *testing.T) {
	// ordered set, value is slice index
	dict := []string{
		"monday",
		"saturday",
		"thursday",
		"tuesday",
	}

	fst := NewFST()

	for value, word := range dict {
		fst.Add(word, value)
	}

	fmt.Print(fst)
}
