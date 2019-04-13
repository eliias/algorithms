/*
	A tree is a binary search tree
*/
package ds

import (
	"fmt"
	"github.com/eliias/algorithms/pkg/strings"
	str "strings"
)

type Edge struct {
	Node   *Node
	Weight int
}

type Node struct {
	Edges map[int]*Edge
	Value []rune
}

type BKTree struct {
	Root              *Node
	calculateDistance func(a []rune, b []rune) int
}

func NewBKTree() *BKTree {
	t := &BKTree{}
	t.calculateDistance = strings.Levenshtein
	return t
}

func (t *BKTree) Insert(value []rune) *Node {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return t.Root
	}

	return t.insert(t.Root, value)
}

func (t *BKTree) insert(node *Node, value []rune) *Node {
	if node.Edges == nil {
		node.Edges = map[int]*Edge{}
	}

	distance := t.calculateDistance(node.Value, value)
	if distance == 0 {
		return node
	}

	edge, ok := node.Edges[distance]

	if !ok {
		newNode := &Node{Value: value}
		node.Edges[distance] = &Edge{
			Weight: distance,
			Node:   newNode,
		}
		return newNode
	} else {
		return t.insert(edge.Node, value)
	}
}

func (t *BKTree) Search(value []rune, tolerance int, limit int) []*Node {
	if limit == 0 {
		limit = 1
	}
	results := []*Node{}
	return t.search(t.Root, value, tolerance, limit, results)
}

func (t *BKTree) search(node *Node, value []rune, tolerance int, limit int, results []*Node) []*Node {
	dist := t.calculateDistance(node.Value, value)
	if dist <= tolerance {
		results = append(results, node)
	}

	for i := dist - tolerance; i <= dist+tolerance; i++ {
		if i < 1 {
			continue
		}
		edge, ok := node.Edges[i]

		if ok {
			results = t.search(edge.Node, value, tolerance, limit, results)
		}
	}

	return results
}

func (t *BKTree) Print() string {
	return t.print("", t.Root, 0, 0)
}

func (t *BKTree) print(s string, node *Node, depth int, w int) string {
	s += fmt.Sprintf(
		str.Repeat("│   ", depth)+"├── (%v) %v\n",
		w, string(node.Value),
	)

	for w, edge := range node.Edges {
		s = t.print(s, edge.Node, depth+1, w)
	}

	return s
}
