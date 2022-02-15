package ds

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree_Insert(t *testing.T) {
	tree, _ := NewBTree(3)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(30)
	tree.Insert(7)
	tree.Insert(17)
}

func TestBTreeNode_Search(t *testing.T) {
	tree, _ := NewBTree(3)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(30)
	tree.Insert(7)
	tree.Insert(17)

	assert.Nil(t, tree.Search(15))
	assert.NotNil(t, tree.Search(6))
	n := tree.Search(31)
	fmt.Print(n)
}

func BenchmarkBTree_Insert(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"deg(2)", 2},
		{"deg(4)", 4},
		{"deg(8)", 8},
		{"deg(16)", 16},
		{"deg(32)", 32}, // probably a good production setting
		{"deg(64)", 64}, // probably a good production setting
		{"deg(100)", 100},
		{"deg(1k)", 1024 * 1},
		{"deg(2k)", 1024 * 2},
		{"deg(4k)", 1024 * 4},
		{"deg(8k)", 1024 * 8},
		{"deg(16k)", 1024 * 16},
	}

	for _, bm := range benchmarks {
		t, _ := NewBTree(bm.n)
		rand.Seed(42)

		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				rand.Seed(42)
				for z := 0; z < 1000; z++ {
					t.Insert(rand.Int())
				}
			}
		})
	}
}

func curve(x float64) float64 {
	// p := (math.Cos(math.Pi * x - math.Pi) + 1) / 2
	return x
}

func BenchmarkBTree_Search(b *testing.B) {
	min := 1
	max := 10_000_000
	steps := 1

	benchmarks := []struct {
		name string
		d    int
	}{
		// {fmt.Sprintf("n=(%v…%v), deg(2)", min, max), 2},
		// {fmt.Sprintf("n=(%v…%v), deg(4)", min, max), 4},
		{fmt.Sprintf("n=(%v…%v), deg(8)", min, max), 8},
		// {fmt.Sprintf("n=(%v…%v), deg(16)", min, max), 16},
		// {fmt.Sprintf("n=(%v…%v), deg(32)", min, max), 32},
		// {fmt.Sprintf("n=(%v…%v), deg(64)", min, max), 64},
		// {fmt.Sprintf("n=(%v…%v), deg(1k)", min, max), 1024 * 1},
		// {fmt.Sprintf("n=(%v…%v), deg(2k)", min, max), 1024 * 2},
		// {fmt.Sprintf("n=(%v…%v), deg(4k)", min, max), 1024 * 4},
		// {fmt.Sprintf("n=(%v…%v), deg(8k)", min, max), 1024 * 8},
	}

	for _, bm := range benchmarks {
		for step := 1; step <= steps; step++ {
			n := int(curve(float64(step)/float64(steps)) * float64(max))
			rand.Seed(42)

			var set []int

			// fill tree
			t, _ := NewBTree(bm.d)
			for i := 0; i < n; i++ {
				rnd := rand.Int()
				set = append(set, rnd)
				t.Insert(rnd)
			}

			len := len(set)

			b.ResetTimer()
			b.Run(fmt.Sprintf("%s, n=%v", bm.name, n), func(b *testing.B) {
				b.ReportAllocs()
				rand.Seed(42)

				for i := 0; i < b.N; i++ {
					t.Search(set[rand.Intn(len-1)])
				}
			})
		}
	}
}
