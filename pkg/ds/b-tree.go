package ds

import (
	"fmt"
)

type BTreeKey = int

type BTree struct {
	Root *BTreeNode
	T    int
}

func NewBTree(t int) (*BTree, error) {
	if t < 2 {
		return nil, fmt.Errorf("minimum degree of b-tree must be > 2, is %v", t)
	}
	return &BTree{nil, t}, nil
}

func (t *BTree) Insert(k int) {
	if t.Root == nil {
		t.Root = NewBTreeNode(true, t.T)
		t.Root.Keys[0] = k
		t.Root.N = 1
		return
	}

	if t.Root.N == 2*t.T-1 {
		s := NewBTreeNode(false, t.T)

		s.Children[0] = t.Root

		s.splitChild(0, t.Root)

		i := 0
		if s.Keys[0] < k {
			i++
			s.Children[i].insertNonFull(k)
		}

		t.Root = s
	} else {
		t.Root.insertNonFull(k)
	}
}

func (t *BTree) Search(k int) *BTreeNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.Search(k)
}

func (t *BTree) Traverse() {
	if t.Root == nil {
		return
	}

	t.Root.Traverse()
}

type BTreeNode struct {
	Children []*BTreeNode
	Keys     []BTreeKey
	Leaf     bool
	N        int
	T        int // minimum degrees
}

func NewBTreeNode(leaf bool, t int) *BTreeNode {
	return &BTreeNode{
		Children: make([]*BTreeNode, 2*t),
		Keys:     make([]BTreeKey, 2*t-1),
		Leaf:     leaf,
		N:        0,
		T:        t,
	}
}

func (n *BTreeNode) Search(k int) *BTreeNode {
	i := 0
	for i < n.N && k > n.Keys[i] {
		i++
	}

	// if n.Keys[i] == k {
	if i < n.N && n.Keys[i] == k {
		return n
	}

	if n.Leaf {
		return nil
	}

	return n.Children[i].Search(k)
}

func (n *BTreeNode) Traverse() {
	// There are n keys and n+1 children, travers through n keys
	// and first n children
	var i int
	for i = 0; i < n.N; i++ {
		// If this is not leaf, then before printing key[i],
		// traverse the subtree rooted with child C[i].
		if !n.Leaf {
			n.Children[i].Traverse()
		}

		fmt.Println(" ", n.Keys[i])
	}

	// Print the subtree rooted with last child
	if !n.Leaf {
		n.Children[i].Traverse()
	}
}

func (n *BTreeNode) splitChild(i int, c *BTreeNode) {
	// Create a new node which is going to store (t-1) keys
	// of y
	z := NewBTreeNode(c.Leaf, c.T)
	z.N = n.T - 1

	// Copy the last (t-1) keys of y to z
	for j := 0; j < n.T-1; j++ {
		z.Keys[j] = c.Keys[j+n.T]
	}

	// Copy the last t children of y to z
	if !c.Leaf {
		for j := 0; j < n.T; j++ {
			z.Children[j] = c.Children[j+n.T]
		}
	}

	// Reduce the number of keys in y
	c.N = n.T - 1

	// Since this node is going to have a new child,
	// create space of new child
	for j := n.N; j >= i+1; j-- {
		n.Children[j+1] = n.Children[j]
	}

	// Link the new child to this node
	n.Children[i+1] = z

	// A key of y will move to this node. Find the location of
	// new key and move all greater keys one space ahead
	for j := n.N - 1; j >= i; j-- {
		n.Keys[j+1] = n.Keys[j]
	}

	// Copy the middle key of y to this node
	n.Keys[i] = c.Keys[n.T-1]

	// Increment count of keys in this node
	n.N++
}

func (n *BTreeNode) insertNonFull(k int) {
	i := n.N - 1

	if n.Leaf {
		// The following loop does two things
		// a) Finds the location of new key to be inserted
		// b) Moves all greater keys to one place ahead
		for i >= 0 && n.Keys[i] > k {
			n.Keys[i+1] = n.Keys[i]
			i--
		}

		// Insert the new key at found location
		n.Keys[i+1] = k
		n.N++
		return
	}

	// If this node is not leaf

	// Find the child which is going to have the new key
	for i >= 0 && n.Keys[i] > k {
		i--
	}

	// See if the found child is full
	if n.Children[i+1].N == 2*n.T-1 {
		// If the child is full, then split it
		n.splitChild(i+1, n.Children[i+1])

		// After split, the middle key of C[i] goes up and
		// C[i] is split into two.  See which of the two
		// is going to have the new key
		if n.Keys[i+1] < k {
			i++
		}
	}

	n.Children[i+1].insertNonFull(k)
}
