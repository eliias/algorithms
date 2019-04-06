// Package math implements a basic library for linear algebra
package math

import (
	"math"
)

// Vector is an alias of type []float64
type Vector []float64

// NullVector of given size
func NullVector(size int) Vector {
	return make(Vector, size, size)
}

func (a Vector) Dot(b Vector) (v float64) {
	if len(a) != len(b) {
		panic("vector: unequal length")
	}

	for index := 0; index < len(a); index++ {
		v += a[index] * b[index]
	}
	return
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Cosine(b Vector) float64 {
	return a.Dot(b) / (a.Length() * b.Length())
}
