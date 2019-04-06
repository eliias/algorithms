package math

import "errors"

type Vector []float64

func (a Vector) Dot(b Vector) (v float64, err error) {
	if len(a) != len(b) {
		err = errors.New("vector: unequal length")
		return
	}
	v = 0
	for index := 0; index < len(a); index++ {
		v += a[index] * b[index]
	}
	return
}
