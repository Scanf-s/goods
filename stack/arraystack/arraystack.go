// Package arraystack implements a stack using a dynamic array (slice).
package arraystack

import (
	"fmt"

	"goods/stack"
)

// ArrayStack is a LIFO stack implemented using a dynamic array.
// It provides O(1) average time for push and pop operations.
type ArrayStack[T any] struct {
	data []T
}

// Compile-time check to ensure ArrayStack implements stack.Stack interface.
var _ stack.Stack[int] = (*ArrayStack[int])(nil)

// New creates and returns an empty ArrayStack.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func New[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{data: make([]T, 0)}
}

// Push adds an element to the top of the stack.
//
// Time Complexity: O(1) average, O(n) worst case when reallocation occurs
// Space Complexity: O(1) average
func (as *ArrayStack[T]) Push(data T) {
	as.data = append(as.data, data)
}

// Pop removes and returns the top element from the stack.
// Returns an error if the stack is empty.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (as *ArrayStack[T]) Pop() (T, error) {
	var defaultValue T
	if len(as.data) == 0 {
		return defaultValue, fmt.Errorf("cannot pop element from empty stack")
	}

	value := as.data[len(as.data)-1]
	as.data = as.data[:len(as.data)-1]
	return value, nil
}

// Top returns the top element without removing it.
// Returns an error if the stack is empty.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (as *ArrayStack[T]) Top() (T, error) {
	var defaultValue T
	if len(as.data) == 0 {
		return defaultValue, fmt.Errorf("cannot get top element from empty stack")
	}
	return as.data[len(as.data)-1], nil
}

// IsEmpty returns true if the stack contains no elements.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (as *ArrayStack[T]) IsEmpty() bool {
	return len(as.data) == 0
}

// Size returns the number of elements in the stack.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (as *ArrayStack[T]) Size() int {
	return len(as.data)
}
