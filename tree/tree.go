package tree

type (

	// Node represents a single node in tree
	Node[T comparable] struct {
		Data T

		Parent *Node[T]

		Children []*Node[T]
	}

	Tree[T comparable] interface {
		IsEmpty() bool

		Clear()

		Contains(element T) bool

		Get(element T) (*Node[T], error)

		Height() int

		BreadthFirstSearch() ([]T, error)

		DepthFirstSearch() ([]T, error)
	}
)

func (n *Node[T]) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node[T]) IsLeaf() bool {
	return len(n.Children) == 0
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
