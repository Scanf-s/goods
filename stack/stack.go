// Package stack provides stack data structure interfaces and implementations.
package stack

// Stack is a LIFO (Last-In-First-Out) data structure interface.
// Elements are added and removed from the same end (top).
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(element T)

	// Pop removes and returns the top element from the stack.
	// Returns an error if the stack is empty.
	Pop() (T, error)

	// Top returns the top element without removing it.
	// Returns an error if the stack is empty.
	Top() (T, error)

	// IsEmpty returns true if the stack contains no elements.
	IsEmpty() bool

	// Size returns the number of elements in the stack.
	Size() int
}
