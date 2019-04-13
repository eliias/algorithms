package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeftMismatch(t *testing.T) {
	a := []rune("ab")
	b := []rune("aa")

	s1, s2 := LeftMismatch(a, 0, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 1, s1, "mismatch left index of first string is 1")
	assert.Equal(t, 1, s2, "mismatch left index of first string is 1")
}

func TestLeftMismatch2(t *testing.T) {
	a := []rune("aaaa")
	b := []rune("aaaa")

	s1, s2 := LeftMismatch(a, 0, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 3, s1, "mismatch left index of first string is 1")
	assert.Equal(t, 3, s2, "mismatch left index of first string is 1")
}

func TestLeftMismatchWithOffset(t *testing.T) {
	a := []rune("ab")
	b := []rune("aa")

	s1, s2 := LeftMismatch(a, 1, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 1, s1, "mismatch left index of first string is 1")
	assert.Equal(t, 0, s2, "mismatch left index of first string is 0")
}

func TestRightMismatch(t *testing.T) {
	a := []rune("ab")
	b := []rune("aa")

	s1, s2 := RightMismatch(a, 0, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 0, s1, "mismatch right index of first string is 0")
	assert.Equal(t, 0, s2, "mismatch right index of first string is 0")
}

func TestRightMismatch2(t *testing.T) {
	a := []rune("abaa")
	b := []rune("aaaa")

	s1, s2 := RightMismatch(a, 0, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 2, s1, "mismatch right index of first string is 0")
	assert.Equal(t, 2, s2, "mismatch right index of first string is 0")
}

func TestRightMismatchWithOffset(t *testing.T) {
	a := []rune("ba")
	b := []rune("aa")

	s1, s2 := LeftMismatch(a, 1, len(a)-1, b, 0, len(b)-1)
	assert.Equal(t, 1, s1, "mismatch right index of first string is 1")
	assert.Equal(t, 0, s2, "mismatch right index of first string is 0")
}

func BenchmarkTestLeftMismatch(b *testing.B) {
	s1 := []rune("aa")
	s2 := []rune("ab")

	for n := 0; n < b.N; n++ {
		LeftMismatch(s1, 0, len(s1)-1, s2, 0, len(s2)-1)
	}
}

func BenchmarkTestLeftMismatch2(b *testing.B) {
	s1 := []rune("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	s2 := []rune("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab")

	for n := 0; n < b.N; n++ {
		LeftMismatch(s1, 0, len(s1)-1, s2, 0, len(s2)-1)
	}
}

func BenchmarkTestRightMismatch(b *testing.B) {
	s1 := []rune("aa")
	s2 := []rune("ba")

	for n := 0; n < b.N; n++ {
		RightMismatch(s1, 0, len(s1)-1, s2, 0, len(s2)-1)
	}
}

func BenchmarkTestRightMismatch2(b *testing.B) {
	s1 := []rune("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	s2 := []rune("baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	for n := 0; n < b.N; n++ {
		RightMismatch(s1, 0, len(s1)-1, s2, 0, len(s2)-1)
	}
}
