// Package math implements a basic library for linear algebra
package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Add(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	sum := Matrix{
		{0, 2, 4},
		{6, 8, 10},
	}

	assert.Equal(t, sum, m1.Add(m2), "addition of two matrices")
}

func TestMatrix_ScalarMultiply(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	v := 2.0

	product := Matrix{
		{0, 2, 4},
		{6, 8, 10},
	}

	assert.Equal(t, product, m1.ScalarMultiply(v), "scalar multiplication of matrix")
}

func TestMatrix_Multiply(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 1},
		{2, 3},
		{4, 5},
	}

	product := Matrix{
		{0*0 + 1*2 + 2*4, 0*1 + 1*3 + 2*5},
		{3*0 + 4*2 + 5*4, 3*1 + 4*3 + 5*5},
	}

	assert.Equal(t, product, m1.Multiply(m2), "multiply of matrices")
}

func TestMatrix_Multiply2(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	assert.Panics(t, func() {
		m1.Multiply(m2)
	})
}

func TestMatrix_M(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 1},
		{2, 3},
		{4, 5},
	}

	assert.Equal(t, 2, m1.M(), "returns number of rows in matrix")
	assert.Equal(t, 3, m2.M(), "returns number of rows in matrix")
}

func TestMatrix_N(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 1},
		{2, 3},
		{4, 5},
	}

	assert.Equal(t, 3, m1.N(), "returns number of columns in matrix")
	assert.Equal(t, 2, m2.N(), "returns number of columns in matrix")
}

func TestMatrix_Transpose(t *testing.T) {
	m1 := Matrix{
		{0, 1, 2},
		{3, 4, 5},
	}

	m2 := Matrix{
		{0, 3},
		{1, 4},
		{2, 5},
	}

	assert.Equal(t, m2, m1.Transpose(), "transpose given matrix")
}
