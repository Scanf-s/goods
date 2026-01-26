package linkedlist

import (
	"fmt"
	"goods/list"
)

// SinglyLinkedList represents a singly linked list data structure.
type SinglyLinkedList[T any] struct {
	// head points to the first node in the list
	head *list.Node[T]

	// tail points to the last node in the list
	tail *list.Node[T]

	// nodeCount represents the number of nodes currently in the list
	nodeCount int
}

// NewSinglyLinkedList returns an empty SinglyLinkedList.
// Time Complexity: O(1)
// Space Complexity: O(1)
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{head: nil, tail: nil, nodeCount: 0}
}

// Append adds an element to the end of the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Append(data T) error {
	if sl.head == nil {
		sl.head = list.NewNode(data)
		sl.tail = sl.head
		sl.nodeCount++
		return nil
	}

	sl.tail.Next = list.NewNode(data)
	sl.tail = sl.tail.Next
	sl.nodeCount++
	return nil
}

// AppendAll adds multiple elements to the end of the list.
// Time Complexity: O(k) where k is number of elements
// Space Complexity: O(k)
func (sl *SinglyLinkedList[T]) AppendAll(data ...T) error {
	for _, element := range data {
		if err := sl.Append(element); err != nil {
			return err
		}
	}
	return nil
}

// Add inserts an element at the specified index, shifting later elements.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Add(index int, data T) error {
	if err := sl.checkIndexRange(index); err != nil {
		return err
	}

	// If index refers to the head of the sll
	if index == 0 {
		node := list.NewNode(data)
		node.Next = sl.head
		sl.head = node
		sl.nodeCount++
		return nil
	}

	// Move to the position and insert a new node
	currentNode := sl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	node := list.NewNode(data)
	node.Next = currentNode.Next
	currentNode.Next = node
	sl.nodeCount++
	return nil
}

// Set replaces the element at the specified index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Set(index int, data T) error {
	if err := sl.checkIndexRange(index); err != nil {
		return err
	}

	currentNode := sl.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Data = data
	return nil
}

// Get returns the element at the specified index.
// If an invalid index is provided, it returns the default value for T type.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Get(index int) (T, error) {
	var defaultValue T
	if err := sl.checkIndexRange(index); err != nil {
		return defaultValue, err
	}

	currentNode := sl.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode.Data, nil
}

// Delete removes the element at the specified index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Delete(index int) error {
	if err := sl.checkIndexRange(index); err != nil {
		return err
	}

	if index == 0 {
		sl.head = sl.head.Next
		if sl.head == nil {
			sl.tail = sl.head
		}
		sl.nodeCount--
		return nil
	}

	currentNode := sl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == nil {
		sl.tail = currentNode
	}
	sl.nodeCount--
	return nil
}

// Size returns the number of elements in the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Size() int {
	return sl.nodeCount
}

// IsEmpty returns true if the list contains no elements.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) IsEmpty() bool {
	return sl.nodeCount == 0
}

// Clear removes all elements from the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Clear() {
	sl.nodeCount = 0
	sl.tail = nil
	sl.head = nil
}

// checkIndexRange validates that index is within [0, nodeCount).
func (sl *SinglyLinkedList[T]) checkIndexRange(index int) error {
	if index < 0 || index >= sl.nodeCount {
		return fmt.Errorf("index %d out of range [0, %d)", index, sl.nodeCount)
	}
	return nil
}
