package binarysearchtree_test

import (
	"slices"
	"testing"
	"time"

	"github.com/Scanf-s/goods/tree"
	bst "github.com/Scanf-s/goods/tree/binary_search_tree"
)

// buildBST inserts values in an order that produces this BST:
//
//	        50
//	      /    \
//	    30      70
//	   /  \    /  \
//	  20  40  60  80
func buildBST(t *testing.T) *bst.BinarySearchTree[int] {
	t.Helper()
	b := bst.NewBinarySearchTree[int]()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		if err := b.Add(v); err != nil {
			t.Fatalf("Add(%d) returned unexpected error: %v", v, err)
		}
	}
	return b
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
	b := bst.NewBinarySearchTree[int]()
	if !b.IsEmpty() {
		t.Error("new tree should be empty")
	}

	if err := b.Add(10); err != nil {
		t.Fatalf("Add(10) returned unexpected error: %v", err)
	}
	if b.IsEmpty() {
		t.Error("tree with one element should not be empty")
	}

	b.Clear()
	if !b.IsEmpty() {
		t.Error("tree should be empty after Clear")
	}
}

func TestAddBuildsCorrectBSTStructure(t *testing.T) {
	b := buildBST(t)

	root := b.Root
	checkNode(t, root, 50, "Root")
	if root == nil {
		return
	}
	if root.Parent != nil {
		t.Error("root's Parent should be nil")
	}

	checkNode(t, root.Left, 30, "Root.Left")
	checkNode(t, root.Right, 70, "Root.Right")
	if root.Left != nil {
		checkNode(t, root.Left.Left, 20, "Root.Left.Left")
		checkNode(t, root.Left.Right, 40, "Root.Left.Right")
		if root.Left.Parent != root {
			t.Error("node 30's Parent should be the root")
		}
	}
	if root.Right != nil {
		checkNode(t, root.Right.Left, 60, "Root.Right.Left")
		checkNode(t, root.Right.Right, 80, "Root.Right.Right")
	}
}

// Duplicates are not rejected: an equal element goes to the right subtree.
func TestAddDuplicate(t *testing.T) {
	b := bst.NewBinarySearchTree[int]()
	if err := b.Add(50); err != nil {
		t.Fatalf("Add(50) returned unexpected error: %v", err)
	}
	if err := b.Add(50); err != nil {
		t.Fatalf("second Add(50) returned unexpected error: %v", err)
	}

	if !b.Contains(50) {
		t.Error("Contains(50) = false after adding duplicates")
	}
	checkNode(t, b.Root.Right, 50, "Root.Right")

	got, err := runWithTimeout(t, "BreadthFirstSearch", b.BreadthFirstSearch)
	if err != nil {
		t.Fatalf("BreadthFirstSearch returned unexpected error: %v", err)
	}
	want := []int{50, 50}
	if !slices.Equal(got, want) {
		t.Errorf("BreadthFirstSearch = %v, want %v (both duplicates kept)", got, want)
	}
}

func TestContains(t *testing.T) {
	b := buildBST(t)

	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		if !b.Contains(v) {
			t.Errorf("Contains(%d) = false, want true", v)
		}
	}
	for _, v := range []int{10, 45, 65, 100} {
		if b.Contains(v) {
			t.Errorf("Contains(%d) = true, want false", v)
		}
	}

	empty := bst.NewBinarySearchTree[int]()
	if empty.Contains(1) {
		t.Error("Contains on empty tree should return false")
	}
}

func TestGet(t *testing.T) {
	b := buildBST(t)

	node, ok := b.Get(40)
	if !ok || node == nil {
		t.Fatal("Get(40) should find the node")
	}
	if node.Data != 40 {
		t.Errorf("Get(40) returned node with data %d", node.Data)
	}
	if node.Parent == nil || node.Parent.Data != 30 {
		t.Errorf("node 40's parent = %v, want 30", node.Parent)
	}
	if !node.IsLeaf() {
		t.Error("node 40 should be a leaf")
	}
	if node.GetLevel() != 2 {
		t.Errorf("node 40's level = %d, want 2", node.GetLevel())
	}

	if _, ok := b.Get(999); ok {
		t.Error("Get(999) should report not found")
	}

	empty := bst.NewBinarySearchTree[int]()
	if _, ok := empty.Get(1); ok {
		t.Error("Get on empty tree should report not found")
	}
}

func TestHeight(t *testing.T) {
	empty := bst.NewBinarySearchTree[int]()
	if h := empty.Height(); h != -1 {
		t.Errorf("Height of empty tree = %d, want -1", h)
	}

	single := bst.NewBinarySearchTree[int]()
	_ = single.Add(1)
	if h := single.Height(); h != 0 {
		t.Errorf("Height of single-node tree = %d, want 0", h)
	}

	if h := buildBST(t).Height(); h != 2 {
		t.Errorf("Height of balanced 7-node tree = %d, want 2", h)
	}

	// Ascending input degenerates a BST into a right chain: 3 edges.
	chain := bst.NewBinarySearchTree[int]()
	for _, v := range []int{10, 20, 30, 40} {
		_ = chain.Add(v)
	}
	if h := chain.Height(); h != 3 {
		t.Errorf("Height of 4-node right chain = %d, want 3", h)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	empty := bst.NewBinarySearchTree[int]()
	if _, err := empty.BreadthFirstSearch(); err == nil {
		t.Error("BreadthFirstSearch on empty tree should return an error")
	}

	b := buildBST(t)
	got, err := runWithTimeout(t, "BreadthFirstSearch", b.BreadthFirstSearch)
	if err != nil {
		t.Fatalf("BreadthFirstSearch returned unexpected error: %v", err)
	}
	want := []int{50, 30, 70, 20, 40, 60, 80} // level order
	if !slices.Equal(got, want) {
		t.Errorf("BreadthFirstSearch = %v, want %v", got, want)
	}
}

func TestDepthFirstSearch(t *testing.T) {
	empty := bst.NewBinarySearchTree[int]()
	if _, err := empty.DepthFirstSearch(); err == nil {
		t.Error("DepthFirstSearch on empty tree should return an error")
	}

	b := buildBST(t)
	got, err := runWithTimeout(t, "DepthFirstSearch", b.DepthFirstSearch)
	if err != nil {
		t.Fatalf("DepthFirstSearch returned unexpected error: %v", err)
	}
	// Preorder: visit node, then its full left subtree, then its right subtree.
	want := []int{50, 30, 20, 40, 70, 60, 80}
	if !slices.Equal(got, want) {
		t.Errorf("DepthFirstSearch = %v, want %v", got, want)
	}

	// Node 30 ends up with only a left child: traversal must handle
	// single-child nodes without panicking.
	//
	//	     50
	//	    /  \
	//	  30    70
	//	  /
	//	20
	incomplete := bst.NewBinarySearchTree[int]()
	for _, v := range []int{50, 30, 70, 20} {
		_ = incomplete.Add(v)
	}
	got, err = runWithTimeout(t, "DepthFirstSearch(incomplete)", incomplete.DepthFirstSearch)
	if err != nil {
		t.Fatalf("DepthFirstSearch on incomplete tree returned unexpected error: %v", err)
	}
	want = []int{50, 30, 20, 70}
	if !slices.Equal(got, want) {
		t.Errorf("DepthFirstSearch on incomplete tree = %v, want %v", got, want)
	}
}
