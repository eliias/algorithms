// Package math implements a basic library for linear algebra
package math

import (
	"fmt"
)

// Row is an alias of type []float64
type Row []float64

// Matrix is an alias of type []Row
type Matrix []Row

// Add builds sum of two matrices
func (m Matrix) Add(m2 Matrix) Matrix {
	if m.M() != m2.M() || m.N() != m2.N() {
		panic("matrix size does not match")
	}

	result := make(Matrix, m.M())
	for row, r := range m {
		result[row] = make(Row, m.N())
		for col := range r {
			result[row][col] = m[row][col] + m2[row][col]
		}
	}

	return result
}

// ScalarMultiply multiply each cell with given value
func (m Matrix) ScalarMultiply(v float64) Matrix {
	result := make(Matrix, m.M())

	for row, r := range m {
		result[row] = make(Row, m.N())
		for col := range r {
			result[row][col] = m[row][col] * v
		}
	}

	return result
}

// Multiply two matrices
func (m Matrix) Multiply(m2 Matrix) Matrix {
	if m.N() != m2.M() {
		panic("m.N != m2.M")
	}

	resultM := m.M()
	resultN := m2.N()

	result := make(Matrix, resultM)

	for row := 0; row < resultM; row++ {
		result[row] = make(Row, resultN)
		for col := 0; col < resultN; col++ {
			v := 0.0
			for i := 0; i < m.N(); i++ {
				v1 := m[row][i]
				v2 := m2[i][col]
				v += v1 * v2
			}
			result[row][col] = v
		}
	}

	return result
}

// Transpose the matrix
func (m Matrix) Transpose() Matrix {
	r := make(Matrix, len(m[0]))

	for x := range r {
		r[x] = make(Row, len(m))
	}
	for y, s := range m {
		for x, e := range s {
			r[x][y] = e
		}
	}

	return r
}

// N of m by n matrix
func (m Matrix) N() int {
	return len(m[0])
}

// M of m by n matrix
func (m Matrix) M() int {
	return len(m)
}

// Print human readable representation
func (m Matrix) Print() string {
	str := ""

	for _, row := range m {
		str += "│"
		for index, col := range row {

			str += fmt.Sprintf("%10.5f", col)
			if index < len(row)-1 {
				str += ", "
			}
		}
		str += "│\n"
	}

	return str
}
