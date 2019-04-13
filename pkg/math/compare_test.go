package math

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMinimum(t *testing.T) {
	result := Minimum(math.MinInt64, 0, math.MaxInt64)
	assert.Equal(t, math.MinInt64, result, "return minimum")
}

func TestMinimum2(t *testing.T) {
	result := Minimum(0, 1, math.MaxInt64)
	assert.Equal(t, 0, result, "return minimum")
}

func TestMinimum3(t *testing.T) {
	result := Minimum(math.MaxInt64-1, math.MaxInt64)
	assert.Equal(t, math.MaxInt64-1, result, "return minimum")
}

func TestMinimum4(t *testing.T) {
	assert.Panics(t, func() {
		Minimum()
	})
}

func TestMinimum5(t *testing.T) {
	result := Minimum(0)
	assert.Equal(t, 0, result, "return minimum")
}

func TestMaximum(t *testing.T) {
	result := Maximum(math.MinInt64, 0, math.MaxInt64)
	assert.Equal(t, math.MaxInt64, result, "return maximum")
}

func TestMaximum2(t *testing.T) {
	result := Maximum(math.MinInt64, 0)
	assert.Equal(t, 0, result, "return maximum")
}

func TestMaximum3(t *testing.T) {
	result := Maximum(math.MinInt64, math.MinInt64+1)
	assert.Equal(t, math.MinInt64+1, result, "return maximum")
}

func TestMaximum4(t *testing.T) {
	assert.Panics(t, func() {
		Maximum()
	})
}

func TestMaximum5(t *testing.T) {
	result := Maximum(0)
	assert.Equal(t, 0, result, "return maximum")
}

func BenchmarkMinimum(b *testing.B) {
	values, _ := faker.RandomInt(math.MinInt8, math.MaxInt8)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Minimum(values...)
	}
}

func BenchmarkMinimum2(b *testing.B) {
	values, _ := faker.RandomInt(math.MinInt16, math.MaxInt16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Minimum(values...)
	}
}

func BenchmarkMaximum(b *testing.B) {
	values, _ := faker.RandomInt(math.MinInt8, math.MaxInt8)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Maximum(values...)
	}
}

func BenchmarkMaximum2(b *testing.B) {
	values, _ := faker.RandomInt(math.MinInt16, math.MaxInt16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Maximum(values...)
	}
}
