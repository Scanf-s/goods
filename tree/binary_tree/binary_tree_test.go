package binarytree_test

import (
	"slices"
	"testing"
	"time"

	"github.com/Scanf-s/goods/tree"
	binarytree "github.com/Scanf-s/goods/tree/binary_tree"
)

func newNode(v int) *tree.Node[int] {
	return &tree.Node[int]{Data: v}
}

func link(parent, left, right *tree.Node[int]) {
	parent.Left = left
	parent.Right = right
	if left != nil {
		left.Parent = parent
	}
	if right != nil {
		right.Parent = parent
	}
}

// buildManualTree wires nodes directly (bypassing Add) so that Get/Height/
// traversal tests don't depend on Add's placement logic. The values are
// deliberately NOT in BST order — this is a plain binary tree:
//
//	        1
//	      /   \
//	     2     3
//	    / \   / \
//	   4   5 6   7
func buildManualTree() *binarytree.BinaryTree[int] {
	n1, n2, n3 := newNode(1), newNode(2), newNode(3)
	n4, n5, n6, n7 := newNode(4), newNode(5), newNode(6), newNode(7)
	link(n1, n2, n3)
	link(n2, n4, n5)
	link(n3, n6, n7)

	bt := binarytree.NewBinaryTree[int]()
	bt.Root = n1
	return bt
}

// checkNode reports an error if node is nil or holds unexpected data.
func checkNode(t *testing.T, node *tree.Node[int], want int, path string) {
	t.Helper()
	if node == nil {
		t.Errorf("%s is nil, want node with data %d", path, want)
		return
	}
	if node.Data != want {
		t.Errorf("%s.Data = %d, want %d", path, node.Data, want)
	}
}

// buildIncompleteTree wires a tree where node 2 has only a left child, so
// traversals must handle single-child nodes:
//
//	     1
//	    / \
//	   2   3
//	  /
//	 4
func buildIncompleteTree() *binarytree.BinaryTree[int] {
	n1, n2, n3, n4 := newNode(1), newNode(2), newNode(3), newNode(4)
	link(n1, n2, n3)
	link(n2, n4, nil)

	bt := binarytree.NewBinaryTree[int]()
	bt.Root = n1
	return bt
}

// runWithTimeout fails the test instead of letting `go test` hang if a
// traversal never terminates, and converts a panic inside the traversal
// into a test failure instead of crashing the test binary.
func runWithTimeout(t *testing.T, name string, fn func() ([]int, error)) ([]int, error) {
	t.Helper()
	type result struct {
		values   []int
		err      error
		panicked any
	}
	done := make(chan result, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- result{panicked: r}
			}
		}()
		values, err := fn()
		done <- result{values: values, err: err}
	}()
	select {
	case r := <-done:
		if r.panicked != nil {
			t.Fatalf("%s panicked: %v", name, r.panicked)
		}
		return r.values, r.err
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not finish within 2s — possible infinite loop", name)
		return nil, nil
	}
}

func TestIsEmptyAndClear(t *testing.T) {
	bt := binarytree.NewBinaryTree[int]()
	if !bt.IsEmpty() {
		t.Error("new tree should be empty")
	}

	if err := bt.Add(10); err != nil {
		t.Fatalf("Add(10) returned unexpected error: %v", err)
	}
	if bt.IsEmpty() {
		t.Error("tree with one element should not be empty")
	}

	bt.Clear()
	if !bt.IsEmpty() {
		t.Error("tree should be empty after Clear")
	}
}

// A plain (non-BST) binary tree conventionally fills level by level, left to
// right, keeping the tree complete. Adding 1..7 should produce:
//
//	        1
//	      /   \
//	     2     3
//	    / \   / \
//	   4   5 6   7
func TestAddFillsLevelOrder(t *testing.T) {
	bt := binarytree.NewBinaryTree[int]()
	for v := 1; v <= 7; v++ {
		if err := bt.Add(v); err != nil {
			t.Fatalf("Add(%d) returned unexpected error: %v", v, err)
		}
	}

	root := bt.Root
	checkNode(t, root, 1, "Root")
	if root == nil {
		return
	}
	if root.Parent != nil {
		t.Error("root's Parent should be nil")
	}

	checkNode(t, root.Left, 2, "Root.Left")
	checkNode(t, root.Right, 3, "Root.Right")
	if root.Left != nil {
		checkNode(t, root.Left.Left, 4, "Root.Left.Left")
		checkNode(t, root.Left.Right, 5, "Root.Left.Right")
		if root.Left.Parent != root {
			t.Error("node 2's Parent should be the root")
		}
	}
	if root.Right != nil {
		checkNode(t, root.Right.Left, 6, "Root.Right.Left")
		checkNode(t, root.Right.Right, 7, "Root.Right.Right")
	}
}

func TestContains(t *testing.T) {
	bt := binarytree.NewBinaryTree[int]()
	values := []int{5, 3, 8, 1, 9, 2, 7}
	for _, v := range values {
		if err := bt.Add(v); err != nil {
			t.Fatalf("Add(%d) returned unexpected error: %v", v, err)
		}
	}

	for _, v := range values {
		if !bt.Contains(v) {
			t.Errorf("Contains(%d) = false, want true", v)
		}
	}
	for _, v := range []int{4, 6, 100} {
		if bt.Contains(v) {
			t.Errorf("Contains(%d) = true, want false", v)
		}
	}

	empty := binarytree.NewBinaryTree[int]()
	if empty.Contains(1) {
		t.Error("Contains on empty tree should return false")
	}
}

// Get must find every element of a plain binary tree, where node placement
// carries no ordering information.
func TestGet(t *testing.T) {
	bt := buildManualTree()

	for v := 1; v <= 7; v++ {
		node, ok := bt.Get(v)
		if !ok || node == nil {
			t.Errorf("Get(%d) should find the node", v)
			continue
		}
		if node.Data != v {
			t.Errorf("Get(%d) returned node with data %d", v, node.Data)
		}
	}

	node, ok := bt.Get(4)
	if ok && node != nil {
		if node.Parent == nil || node.Parent.Data != 2 {
			t.Errorf("node 4's parent = %v, want 2", node.Parent)
		}
		if !node.IsLeaf() {
			t.Error("node 4 should be a leaf")
		}
		if node.GetLevel() != 2 {
			t.Errorf("node 4's level = %d, want 2", node.GetLevel())
		}
	}

	if _, ok := bt.Get(999); ok {
		t.Error("Get(999) should report not found")
	}
}

func TestHeight(t *testing.T) {
	empty := binarytree.NewBinaryTree[int]()
	if h := empty.Height(); h != -1 {
		t.Errorf("Height of empty tree = %d, want -1", h)
	}

	single := binarytree.NewBinaryTree[int]()
	single.Root = newNode(1)
	if h := single.Height(); h != 0 {
		t.Errorf("Height of single-node tree = %d, want 0", h)
	}

	if h := buildManualTree().Height(); h != 2 {
		t.Errorf("Height of perfect 7-node tree = %d, want 2", h)
	}

	// Left-skewed chain 10 -> 20 -> 30 -> 40: 3 edges from root to deepest leaf.
	chain := binarytree.NewBinaryTree[int]()
	n10, n20, n30, n40 := newNode(10), newNode(20), newNode(30), newNode(40)
	link(n10, n20, nil)
	link(n20, n30, nil)
	link(n30, n40, nil)
	chain.Root = n10
	if h := chain.Height(); h != 3 {
		t.Errorf("Height of 4-node chain = %d, want 3", h)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	empty := binarytree.NewBinaryTree[int]()
	if _, err := empty.BreadthFirstSearch(); err == nil {
		t.Error("BreadthFirstSearch on empty tree should return an error")
	}

	bt := buildManualTree()
	got, err := runWithTimeout(t, "BreadthFirstSearch", bt.BreadthFirstSearch)
	if err != nil {
		t.Fatalf("BreadthFirstSearch returned unexpected error: %v", err)
	}
	want := []int{1, 2, 3, 4, 5, 6, 7} // level order
	if !slices.Equal(got, want) {
		t.Errorf("BreadthFirstSearch = %v, want %v", got, want)
	}

	incomplete := buildIncompleteTree()
	got, err = runWithTimeout(t, "BreadthFirstSearch(incomplete)", incomplete.BreadthFirstSearch)
	if err != nil {
		t.Fatalf("BreadthFirstSearch on incomplete tree returned unexpected error: %v", err)
	}
	want = []int{1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Errorf("BreadthFirstSearch on incomplete tree = %v, want %v", got, want)
	}
}

func TestDepthFirstSearch(t *testing.T) {
	empty := binarytree.NewBinaryTree[int]()
	if _, err := empty.DepthFirstSearch(); err == nil {
		t.Error("DepthFirstSearch on empty tree should return an error")
	}

	bt := buildManualTree()
	got, err := runWithTimeout(t, "DepthFirstSearch", bt.DepthFirstSearch)
	if err != nil {
		t.Fatalf("DepthFirstSearch returned unexpected error: %v", err)
	}
	// Preorder: visit node, then its full left subtree, then its right subtree.
	want := []int{1, 2, 4, 5, 3, 6, 7}
	if !slices.Equal(got, want) {
		t.Errorf("DepthFirstSearch = %v, want %v", got, want)
	}

	incomplete := buildIncompleteTree()
	got, err = runWithTimeout(t, "DepthFirstSearch(incomplete)", incomplete.DepthFirstSearch)
	if err != nil {
		t.Fatalf("DepthFirstSearch on incomplete tree returned unexpected error: %v", err)
	}
	want = []int{1, 2, 4, 3}
	if !slices.Equal(got, want) {
		t.Errorf("DepthFirstSearch on incomplete tree = %v, want %v", got, want)
	}
}
