package linkedlist

import (
	"fmt"
	"goods/list"
)

// DoublyLinkedList represents a doubly linked list data structure.
type DoublyLinkedList[T any] struct {
	// head points to the first node in the list
	head *list.Node[T]

	// tail points to the last node in the list
	tail *list.Node[T]

	// nodeCount represents the number of nodes currently in the list
	nodeCount int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{head: nil, tail: nil, nodeCount: 0}
}

func (dl *DoublyLinkedList[T]) Append(data T) error {
	if dl.head == nil {
		dl.head = list.NewNode(data)
		dl.tail = dl.head
		dl.nodeCount++
		return nil
	}

	newNode := list.NewNode(data)
	dl.tail.Next = newNode
	newNode.Prev = dl.tail
	dl.tail = newNode
	dl.nodeCount++
	return nil
}

func (dl *DoublyLinkedList[T]) AppendAll(data ...T) error {
	for _, element := range data {
		if err := dl.Append(element); err != nil {
			return err
		}
	}
	return nil
}

func (dl *DoublyLinkedList[T]) Add(index int, data T) error {
	if err := dl.checkIndexRange(index); err != nil {
		return err
	}

	// Prepare new node
	newNode := list.NewNode(data)

	// If index refers to the head of the list
	if index == 0 {
		newNode.Next = dl.head
		dl.head.Prev = newNode
		dl.head = newNode
		dl.nodeCount++
		return nil
	}

	currentNode := dl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next.Prev = newNode
	currentNode.Next = newNode
	newNode.Prev = currentNode
	dl.nodeCount++
	return nil
}

func (dl *DoublyLinkedList[T]) Set(index int, data T) error {
	if err := dl.checkIndexRange(index); err != nil {
		return err
	}

	currentNode := dl.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Data = data
	return nil
}

func (dl *DoublyLinkedList[T]) Get(index int) (T, error) {
	var defaultValue T
	if err := dl.checkIndexRange(index); err != nil {
		return defaultValue, err
	}

	currentNode := dl.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode.Data, nil
}

func (dl *DoublyLinkedList[T]) Delete(index int) error {
	if err := dl.checkIndexRange(index); err != nil {
		return err
	}

	if index == 0 {
		dl.head = dl.head.Next
		if dl.head == nil {
			dl.tail = dl.head
			dl.nodeCount--
			return nil
		}
		dl.head.Prev = nil
		dl.nodeCount--
		return nil
	}

	currentNode := dl.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == nil {
		dl.tail = currentNode
	} else {
		currentNode.Next.Prev = currentNode
	}
	dl.nodeCount--
	return nil
}

func (dl *DoublyLinkedList[T]) Size() int {
	return dl.nodeCount
}

func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.nodeCount == 0
}

func (dl *DoublyLinkedList[T]) Clear() {
	dl.head = nil
	dl.tail = nil
	dl.nodeCount = 0
}

// checkIndexRange validates that index is within [0, nodeCount).
func (dl *DoublyLinkedList[T]) checkIndexRange(index int) error {
	if index < 0 || index >= dl.nodeCount {
		return fmt.Errorf("index %d out of range [0, %d)", index, dl.nodeCount)
	}
	return nil
}
