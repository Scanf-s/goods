package linkedlist

import (
	"fmt"

	"github.com/Scanf-s/goods/list"
)

// CircularLinkedList represents a circular singly linked list where the tail
// node's Next points back to the head, forming a ring.
type CircularLinkedList[T any] struct {
	// head points to the first node in the ring
	head *list.Node[T]

	// tail points to the last node in the ring (tail.Next == head)
	tail *list.Node[T]

	// nodeCount represents the number of nodes currently in the list
	nodeCount int
}

// Compile time interface implementation check
var _ list.List[int] = (*CircularLinkedList[int])(nil)

// NewCircularLinkedList returns an empty CircularLinkedList.
// Time Complexity: O(1)
// Space Complexity: O(1)
func NewCircularLinkedList[T any]() *CircularLinkedList[T] {
	return &CircularLinkedList[T]{head: nil, tail: nil, nodeCount: 0}
}

// Append adds an element to the end of the list, keeping the ring closed.
// Time Complexity: O(1)
// Space Complexity: O(1)
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

// AppendAll adds multiple elements to the end of the list.
// Time Complexity: O(k) where k is number of elements
// Space Complexity: O(k)
func (cl *CircularLinkedList[T]) AppendAll(data ...T) error {
	for _, element := range data {
		if err := cl.Append(element); err != nil {
			return err
		}
	}
	return nil
}

// Prepend adds an element in front of the list, keeping the ring closed.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Prepend(data T) error {
	newNode := list.NewNode(data)

	if cl.head == nil {
		cl.head = newNode
		cl.tail = newNode
		newNode.Next = newNode
		cl.nodeCount++
		return nil
	}

	newNode.Next = cl.head
	cl.tail.Next = newNode
	cl.head = newNode
	cl.nodeCount++
	return nil
}

// Add inserts an element at the specified index, shifting later elements.
// The valid range is [0, Size()]: index 0 prepends, index Size() appends.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Add(index int, data T) error {
	if index < 0 || index > cl.nodeCount {
		return fmt.Errorf("index %d out of range [0, %d]", index, cl.nodeCount)
	}

	if index == 0 {
		return cl.Prepend(data)
	}

	if index == cl.nodeCount {
		return cl.Append(data)
	}

	currentNode := cl.head
	for range index - 1 {
		currentNode = currentNode.Next
	}
	node := list.NewNode(data)
	node.Next = currentNode.Next
	currentNode.Next = node
	cl.nodeCount++
	return nil
}

// Set replaces the element at the specified index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Set(index int, data T) error {
	if err := cl.checkIndexRange(index); err != nil {
		return err
	}

	currentNode := cl.head
	for range index {
		currentNode = currentNode.Next
	}
	currentNode.Data = data
	return nil
}

// Get returns the element at the specified index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Get(index int) (T, error) {
	var defaultValue T
	if err := cl.checkIndexRange(index); err != nil {
		return defaultValue, err
	}

	currentNode := cl.head
	for range index {
		currentNode = currentNode.Next
	}
	return currentNode.Data, nil
}

// Delete removes the element at the specified index, keeping the ring closed.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Delete(index int) error {
	if err := cl.checkIndexRange(index); err != nil {
		return err
	}

	if index == 0 {
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
	for range index - 1 {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == cl.head {
		cl.tail = currentNode
	}
	cl.nodeCount--
	return nil
}

// Size returns the number of elements in the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Size() int {
	return cl.nodeCount
}

// IsEmpty returns true if the list contains no elements.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) IsEmpty() bool {
	return cl.nodeCount == 0
}

// Clear removes all elements from the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (cl *CircularLinkedList[T]) Clear() {
	cl.head = nil
	cl.tail = nil
	cl.nodeCount = 0
}

func (cl *CircularLinkedList[T]) Head() (T, error) {
	var element T
	if cl.head == nil {
		return element, fmt.Errorf("the list is empty")
	}
	return cl.head.Data, nil
}

func (cl *CircularLinkedList[T]) PopHead() (T, error) {
	var element T
	if cl.head == nil {
		return element, fmt.Errorf("the list is empty. Cannot pop the head")
	}
	element = cl.head.Data
	cl.head = cl.head.Next
	cl.tail = cl.head
	cl.nodeCount--
	return element, nil
}

// checkIndexRange validates that index is within [0, nodeCount).
func (cl *CircularLinkedList[T]) checkIndexRange(index int) error {
	if index < 0 || index >= cl.nodeCount {
		return fmt.Errorf("index %d out of range [0, %d)", index, cl.nodeCount)
	}
	return nil
}
