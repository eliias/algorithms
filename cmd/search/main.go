package main

import (
	"fmt"
	"github.com/eliias/algorithms/pkg/math"
	"os"
)

func main() {
	m1 := math.Matrix{
		{2, 3, 4},
		{1, 0, 0},
	}

	m2 := math.Matrix{
		{0, 1000},
		{1, 100},
		{0, 10},
	}

	fmt.Fprint(os.Stdout, m1.Print())
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, m2.Print())
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, m1.Multiply(m2).Print())
}
