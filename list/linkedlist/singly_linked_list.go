package linkedlist

import "fmt"

// Node represents a single node in the linked list.
type Node[T any] struct {
	// data stores the element value
	data T

	// next points to the next node in the list
	next *Node[T]
}

// SinglyLinkedList represents a singly linked list data structure.
type SinglyLinkedList[T any] struct {
	// head points to the first node in the list
	head *Node[T]

	// tail points to the last node in the list
	tail *Node[T]

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
		sl.head = newNode(data)
		sl.tail = sl.head
		sl.nodeCount++
		return nil
	}

	sl.tail.next = newNode(data)
	sl.tail = sl.tail.next
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
		node := newNode(data)
		node.next = sl.head
		sl.head = node
		sl.nodeCount++
		return nil
	}

	// Move to the position and insert a new node
	currentNode := sl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	node := newNode(data)
	node.next = currentNode.next
	currentNode.next = node
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
		currentNode = currentNode.next
	}
	currentNode.data = data
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
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

// Delete removes the element at the specified index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (sl *SinglyLinkedList[T]) Delete(index int) error {
	if err := sl.checkIndexRange(index); err != nil {
		return err
	}

	if index == 0 {
		sl.head = sl.head.next
		if sl.head == nil {
			sl.tail = sl.head
		}
		sl.nodeCount--
		return nil
	}

	currentNode := sl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	if currentNode.next == nil {
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

// newNode creates a new node with the given data.
func newNode[T any](data T) *Node[T] {
	return &Node[T]{data: data, next: nil}
}

// checkIndexRange validates that index is within [0, nodeCount).
func (sl *SinglyLinkedList[T]) checkIndexRange(index int) error {
	if index < 0 || index >= sl.nodeCount {
		return fmt.Errorf("index %d out of range [0, %d)", index, sl.nodeCount)
	}
	return nil
}
