package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotProductOfTwoPerpendicularVectors(t *testing.T) {
	result := Vector{1., 0., 0.}.Dot(Vector{0., 1., 0.})
	assert.Equal(t, 0., result, "dot product of two perpendicular vectors is 0")
}

func TestDotProductOfTwoParallelVectors(t *testing.T) {
	result := Vector{1., 0., 0.}.Dot(Vector{1., 0., 0.})
	assert.Equal(t, 1., result, "dot product of two parallel vectors is 1")
}

func TestDotProductOfUnequalVectorsReturnsError(t *testing.T) {
	assert.Panics(t, func() { Vector{1.}.Dot(Vector{1., 0.}) }, "dot product of two unequal vectors is error")
}
