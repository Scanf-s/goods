package linkedlist

import (
	"fmt"
	"goods/list"
)

type CircularLinkedList[T any] struct {
	head *list.Node[T]

	tail *list.Node[T]

	nodeCount int
}

func NewCircularLinkedList[T any]() *CircularLinkedList[T] {
	return &CircularLinkedList[T]{head: nil, tail: nil, nodeCount: 0}
}

func (cl *CircularLinkedList[T]) Append(data T) error {
	node := list.NewNode(data)
	if cl.head == nil {
		cl.head = node
		cl.tail = node
		node.Next = cl.head
		cl.nodeCount++
		return nil
	}

	cl.tail.Next = node
	cl.tail = node
	node.Next = cl.head
	cl.nodeCount++
	return nil
}

func (cl *CircularLinkedList[T]) AppendAll(data ...T) error {
	for _, element := range data {
		if err := cl.Append(element); err != nil {
			return err
		}
	}
	return nil
}

func (cl *CircularLinkedList[T]) Add(index int, data T) error {
	if cl.IsEmpty() {
		return fmt.Errorf("cannot add element when list is empty")
	}
	convertedIndex := cl.convertIndex(index)

	node := list.NewNode(data)
	if convertedIndex == 0 {
		node.Next = cl.head
		cl.tail.Next = node
		cl.head = node
		cl.nodeCount++
		return nil
	}

	currentNode := cl.head
	for i := 0; i < convertedIndex-1; i++ {
		currentNode = currentNode.Next
	}

	node.Next = currentNode.Next
	currentNode.Next = node
	cl.nodeCount++
	return nil
}

func (cl *CircularLinkedList[T]) Set(index int, data T) error {
	if cl.IsEmpty() {
		return fmt.Errorf("cannot set element when list is empty")
	}
	convertedIndex := cl.convertIndex(index)

	currentNode := cl.head
	for i := 0; i < convertedIndex; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Data = data
	return nil
}

func (cl *CircularLinkedList[T]) Get(index int) (T, error) {
	var defaultValue T
	if cl.IsEmpty() {
		return defaultValue, fmt.Errorf("cannot get element when list is empty")
	}
	convertedIndex := cl.convertIndex(index)

	currentNode := cl.head
	for i := 0; i < convertedIndex; i++ {
		currentNode = currentNode.Next
	}
	return currentNode.Data, nil
}

func (cl *CircularLinkedList[T]) Delete(index int) error {
	if cl.IsEmpty() {
		return fmt.Errorf("cannot delete element when list is empty")
	}
	convertedIndex := cl.convertIndex(index)
	if convertedIndex == 0 {
		if cl.Size() == 1 {
			cl.Clear()
			return nil
		}

		cl.head = cl.head.Next
		cl.tail.Next = cl.head
		cl.nodeCount--
		return nil
	}

	currentNode := cl.head
	for i := 0; i < convertedIndex-1; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == cl.head {
		cl.tail = currentNode
	}
	cl.nodeCount--
	return nil
}

func (cl *CircularLinkedList[T]) Size() int {
	return cl.nodeCount
}

func (cl *CircularLinkedList[T]) IsEmpty() bool {
	return cl.nodeCount == 0
}

func (cl *CircularLinkedList[T]) Clear() {
	cl.head = nil
	cl.tail = nil
	cl.nodeCount = 0
}

// checkIndexRange converts index using operation: index % nodeCount
func (cl *CircularLinkedList[T]) convertIndex(index int) int {
	return ((index % cl.nodeCount) + cl.nodeCount) % cl.nodeCount
}
