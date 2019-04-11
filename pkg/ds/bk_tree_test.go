/*
	A tree is a binary search tree
*/
package ds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/eliias/algorithms/pkg/strings"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestBKTree_InsertFirstValue(t *testing.T) {
	tree := NewBKTree()
	tree.Insert([]rune("hannes"))
	assert.Equal(t, []rune("hannes"), tree.Root.Value, "For empty tree, first added value becomes root")
}

func TestBKTree_InsertTwoValues(t *testing.T) {
	tree := NewBKTree()
	tree.Insert([]rune("hannes"))
	tree.Insert([]rune("moser"))

	assert.Equal(t, []rune("hannes"), tree.Root.Value, "For empty tree, first added value becomes root")
	assert.Equal(t, 1, len(tree.Root.Edges), "Root node should have one edge")
	assert.Equal(t, []rune("moser"), tree.Root.Edges[5].Node.Value, "Root node should have an edge with weight 5")
}

func TestBKTree_InsertSameValue(t *testing.T) {
	tree := NewBKTree()
	tree.Insert([]rune("hannes"))
	tree.Insert([]rune("moser"))

	assert.Equal(t, []rune("hannes"), tree.Root.Value, "For empty tree, first added value becomes root")
	assert.Equal(t, 1, len(tree.Root.Edges), "Root node should have one edge")
	assert.Equal(t, []rune("moser"), tree.Root.Edges[5].Node.Value, "Root node should have an edge with weight 5")
}

func TestBKTree_Search(t *testing.T) {
	tree := NewBKTree()
	tree.Insert([]rune("hannes"))
	tree.Insert([]rune("moser"))
	result := tree.Search([]rune("moser"), 1, 0)

	assert.Equal(t, 1, len(result), "Should return one result")
	assert.Equal(t, []rune("moser"), result[0].Value, "Should find second node")
}

func TestBKTree_SearchMaxDistance(t *testing.T) {
	tree := NewBKTree()
	tree.Insert([]rune("hannes"))
	tree.Insert([]rune("moser"))
	result := tree.Search([]rune("moser"), 5, math.MaxInt64)

	assert.Equal(t, 2, len(result), "Should return all nodes")
}

const count = 100000

func BenchmarkBKTree_Insert(b *testing.B) {
	tree := NewBKTree()
	var values [][]rune
	for i := 0; i < count; i++ {
		values = append(values, []rune(faker.FirstName()+faker.LastName()))
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var i int
		if n > count {
			i = n % count
		}
		tree.Insert(values[i])
	}
}

func BenchmarkBKTree_Search1(b *testing.B) {
	tree := NewBKTree()
	var node *Node
	for i := 0; i < count; i++ {
		node = tree.Insert([]rune(faker.FirstName()))
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tree.Search(node.Value, 1, 1)
	}
}

func BenchmarkBKTree_SearchBruteForce(b *testing.B) {
	var nodes [][]rune
	for i := 0; i < count; i++ {
		nodes = append(nodes, []rune(faker.FirstName()))
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r := rand.Intn(count)
		term := nodes[r]

		for i := 0; i < count; i++ {
			distance := strings.Levenshtein(term, nodes[i])
			if distance <= 1 {
				break
			}
		}
	}
}
