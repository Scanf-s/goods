package tree

import "cmp"

type (

	// Node represents a single node in tree
	Node[T cmp.Ordered] struct {
		Data T

		Parent *Node[T]

		Left *Node[T]

		Right *Node[T]
	}

	Tree[T cmp.Ordered] interface {
		IsEmpty() bool

		Clear()

		Add(element T) error

		Contains(element T) bool

		Get(element T) (*Node[T], bool)

		Height() int

		BreadthFirstSearch() ([]T, error)

		DepthFirstSearch() ([]T, error)
	}
)

func (n *Node[T]) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[T]) GetLevel() int {
	curNode := n
	level := 0
	for curNode.Parent != nil {
		curNode = curNode.Parent
		level++
	}
	return level
}
