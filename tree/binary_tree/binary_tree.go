package binarytree

import (
	"cmp"
	"fmt"

	"github.com/Scanf-s/goods/queue/deque"
	"github.com/Scanf-s/goods/tree"
)

type BinaryTree[T cmp.Ordered] struct {
	Root *tree.Node[T]
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		Root: nil,
	}
}

func (b *BinaryTree[T]) IsEmpty() bool {
	if b.Root == nil {
		return true
	} else {
		return false
	}
}

func (b *BinaryTree[T]) Clear() {
	b.Root = nil
}

func (b *BinaryTree[T]) Add(element T) error {
	if b == nil {
		return fmt.Errorf("please initialize binary tree first")
	}

	newNode := &tree.Node[T]{
		Data: element,
		Parent: nil,
		Left: nil,
		Right: nil,
	}
	if b.Root == nil {
		b.Root = newNode
		return nil
	}

	dq := deque.NewDeque[*tree.Node[T]]()
	dq.Offer(b.Root)
	for !dq.IsEmpty() {
		curNode, err := dq.PollFront()
		if err != nil {
			return fmt.Errorf("cannot poll element from deque while add")
		}
		if curNode.Left == nil {
			curNode.Left = newNode
			newNode.Parent = curNode
			return nil
		}
		dq.Offer(curNode.Left)
		if curNode.Right == nil {
			curNode.Right = newNode
			newNode.Parent = curNode
			return nil
		}
		dq.Offer(curNode.Right)
	}
	return nil
}

func (b *BinaryTree[T]) Contains(element T) bool {
	if b == nil || b.Root == nil {
		return false
	}

	if b.Root.Data == element {
		return true
	}

	dq := deque.NewDeque[*tree.Node[T]]()
	dq.Offer(b.Root)
	for !dq.IsEmpty() {
		curNode, err := dq.PollFront()
		if err != nil {
			return false
		}
		if curNode.Data == element {
			return true
		}
		
		if curNode.Left != nil {
			dq.Offer(curNode.Left)
		}
		if curNode.Right != nil {
			dq.Offer(curNode.Right)
		}
	}
	return false
}

func (b *BinaryTree[T]) Get(element T) (*tree.Node[T], bool) {
	if b == nil || b.Root == nil {
		return nil, false
	}

	dq := deque.NewDeque[*tree.Node[T]]()
	dq.Offer(b.Root)
	for !dq.IsEmpty() {
		curNode, err := dq.PollFront()
		if err != nil {
			return nil, false
		}
		if curNode.Data == element {
			return curNode, true
		}
		if curNode.Left != nil {
			dq.Offer(curNode.Left)
		}
		if curNode.Right != nil {
			dq.Offer(curNode.Right)
		}
	}
	return nil, false
}

func (b *BinaryTree[T]) Height() int {
	if b == nil || b.Root == nil {
		return -1
	}

	left := calculateHeight(b.Root.Left)
	right := calculateHeight(b.Root.Right)
	if left > right {
		return left
	}
	return right
}

func calculateHeight[T cmp.Ordered](node *tree.Node[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := calculateHeight(node.Left)
	rightHeight := calculateHeight(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (b *BinaryTree[T]) BreadthFirstSearch() ([]T, error) {
	if b == nil || b.Root == nil {
		return nil, fmt.Errorf("please initialize binary tree first")
	}
	dq := deque.NewDeque[*tree.Node[T]]()
	result := []T{}
	dq.Offer(b.Root)
	result = append(result, b.Root.Data)
	for !dq.IsEmpty() {
		curNode, err := dq.PollFront()
		if err != nil {
			return nil, fmt.Errorf("cannot poll element from deque while bfs")
		}
		
		if curNode.Left != nil {
			dq.Offer(curNode.Left)
			result = append(result, curNode.Left.Data)
		}
		if curNode.Right != nil {
			dq.Offer(curNode.Right)
			result = append(result, curNode.Right.Data)
		}
	}
	return result, nil
}

func (b *BinaryTree[T]) DepthFirstSearch() ([]T, error) {
	if b == nil || b.Root == nil {
		return nil, fmt.Errorf("please initialize binary tree first")
	}

	result := []T{}
	result = append(result, dfsHelper(b.Root)...)
	return result, nil
}

func dfsHelper[T cmp.Ordered](node *tree.Node[T]) []T {
	if node.Left == nil && node.Right == nil {
		return []T{node.Data}
	}

	arr := []T{node.Data}
	if node.Left != nil {
		arr = append(arr, dfsHelper(node.Left)...)
	}
	if node.Right != nil {
		arr = append(arr, dfsHelper(node.Right)...)
	}
	return arr
}
