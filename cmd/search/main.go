package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/eliias/algorithms/pkg/ds"
	"math"
	"os"
)

func main() {
	count := 100
	tree := ds.NewBKTree()
	var values [][]rune

	for i := 0; i < count; i++ {
		value := []rune(faker.FirstName() + faker.LastName())
		values = append(values, value)
		tree.Insert(value)
	}

	fmt.Fprint(os.Stdout, string(values[count/2]))
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, tree.Print())
	fmt.Fprint(os.Stdout, "\n")

	results := tree.Search(values[count/2], 1, math.MaxInt64)

	fmt.Fprint(os.Stdout, len(results))
}
