// Package math implements a basic library for linear algebra
package math

import (
	"math"
)

// Vector is an alias of type []float64
type Vector []float64

// NullVector of given size
func NullVector(size int) Vector {
	return make(Vector, size)
}

func (a Vector) Dot(b Vector) (v float64) {
	if len(a) != len(b) {
		panic("vector: unequal length")
	}

	for index := 0; index < len(a); index++ {
		v += a[index] * b[index]
	}
	return v
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Cosine(b Vector) float64 {
	return a.Dot(b) / (a.Length() * b.Length())
}

type Vector3 [3]float64

func NullVector3() Vector3 {
	return Vector3{0, 0, 0}
}

func (a Vector3) Dot(b Vector3) (v float64) {
	for index := 0; index < len(a); index++ {
		v += a[index] * b[index]
	}
	return v
}

func (a Vector3) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector3) Cosine(b Vector3) float64 {
	return a.Dot(b) / (a.Length() * b.Length())
}

type Vector2 [2]float64

func NullVector2() Vector2 {
	return Vector2{0, 0}
}

func (a Vector2) Dot(b Vector2) (v float64) {
	for index := 0; index < len(a); index++ {
		v += a[index] * b[index]
	}
	return v
}

func (a Vector2) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector2) Cosine(b Vector2) float64 {
	return a.Dot(b) / (a.Length() * b.Length())
}
