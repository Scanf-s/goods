// Package linkedstack implements a stack using a singly linked list.
package linkedstack

import (
	"goods/list/linkedlist"
	"goods/stack"
)

// LinkedStack is a LIFO stack implemented using a singly linked list.
// It provides O(1) time for push and pop operations.
type LinkedStack[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

// Compile-time check to ensure LinkedStack implements stack.Stack interface.
var _ stack.Stack[int] = (*LinkedStack[int])(nil)

// NewLinkedStack creates and returns an empty LinkedStack.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{list: linkedlist.NewSinglyLinkedList[T]()}
}

// Push adds an element to the top of the stack.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (ls *LinkedStack[T]) Push(data T) error {
	return ls.list.Prepend(data)
}

// Pop removes and returns the top element from the stack.
// Returns an error if the stack is empty.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (ls *LinkedStack[T]) Pop() (T, error) {
	return ls.list.PopHead()
}

// Top returns the top element without removing it.
// Returns an error if the stack is empty.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (ls *LinkedStack[T]) Top() (T, error) {
	return ls.list.Head()
}

// IsEmpty returns true if the stack contains no elements.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (ls *LinkedStack[T]) IsEmpty() bool {
	return ls.list.IsEmpty()
}

// Size returns the number of elements in the stack.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (ls *LinkedStack[T]) Size() int {
	return ls.list.Size()
}
