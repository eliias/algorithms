package main

import (
	"fmt"
	"github.com/eliias/algorithms/pkg/strings"
	"os"
)

func main() {
	a := []rune("bbba")
	b := []rune("bbbb")

	fmt.Fprint(os.Stdout, strings.Levenshtein(a, b))
}
