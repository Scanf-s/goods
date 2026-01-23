package arraylist

import (
	"fmt"
)

// ArrayList represents the ArrayList data structure
type ArrayList[T any] struct {
	// data is the array to store the T type elements
	data []T

	// listCapacity represents the capacity of the list (total number of elements that could be stored)
	listCapacity int64

	// listSize represents the number of elements currently stored in the list
	listSize int64
}

// New returns an empty ArrayList with the specified initial capacity.
// Time Complexity: O(1)
// Space Complexity: O(capacity)
func New[T any](capacity int64) *ArrayList[T] {
	return &ArrayList[T]{
		data:         make([]T, capacity),
		listCapacity: capacity,
		listSize:     0, // empty list
	}
}

// Append adds an element to the end of the list.
// Time Complexity: O(1) average, O(n) when resizing
// Space Complexity: O(1) average, O(n) when resizing
func (al *ArrayList[T]) Append(element T) error {
	if al.listSize >= al.listCapacity {
		if err := al.increaseCapacity(); err != nil {
			return fmt.Errorf("failed to increase capacity: %w", err)
		}
	}
	al.data[al.listSize] = element
	al.listSize++
	return nil
}

// AppendAll adds multiple elements to the end of the list.
// Time Complexity: O(k) where k is number of elements, O(n+k) when resizing
// Space Complexity: O(1) average, O(n+k) when resizing
func (al *ArrayList[T]) AppendAll(elements ...T) error {
	required := al.listSize + int64(len(elements))
	for required > al.listCapacity {
		if err := al.increaseCapacity(); err != nil {
			return fmt.Errorf("failed to increase capacity: %w", err)
		}
	}
	for _, elem := range elements {
		al.data[al.listSize] = elem
		al.listSize++
	}
	return nil
}

// Add inserts an element at the specified index, shifting later elements to right.
// Time Complexity: O(n)
// Space Complexity: O(1) average, O(n) when resizing
func (al *ArrayList[T]) Add(index int64, element T) error {
	// Allow inserting at index 0 to listSize (inclusive)
	if index < 0 || index > al.listSize {
		return fmt.Errorf("index %d out of range [0, %d]", index, al.listSize)
	}

	if al.listSize >= al.listCapacity {
		if err := al.increaseCapacity(); err != nil {
			return fmt.Errorf("failed to increase capacity: %w", err)
		}
	}

	// Shift elements to the right (start from end, move backward)
	for i := al.listSize; i > index; i-- {
		al.data[i] = al.data[i-1]
	}

	// Insert the element at the specified index
	al.data[index] = element
	al.listSize++
	return nil
}

// Set replaces the element at the specified index.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (al *ArrayList[T]) Set(index int64, element T) error {
	if err := al.checkIndexRange(index); err != nil {
		return err
	}
	al.data[index] = element
	return nil
}

// Get returns the element at the specified index.
// If an invalid index is provided, it returns the default value for T type.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (al *ArrayList[T]) Get(index int64) (T, error) {
	var defaultValue T
	if err := al.checkIndexRange(index); err != nil {
		return defaultValue, err
	}
	return al.data[index], nil
}

// Delete removes the element at the specified index, shifting later elements left.
// Time Complexity: O(n)
// Space Complexity: O(1) average, O(n) when resizing
func (al *ArrayList[T]) Delete(index int64) error {
	if err := al.checkIndexRange(index); err != nil {
		return err
	}

	// Shift elements to the left
	for i := index; i < al.listSize-1; i++ {
		al.data[i] = al.data[i+1]
	}

	// Clear the last element and decrement size
	var defaultValue T
	al.data[al.listSize-1] = defaultValue
	al.listSize--

	// Reduce capacity if the size is less than 1/2 of capacity
	if al.listSize > 0 && (al.listSize*2) <= al.listCapacity {
		if err := al.decreaseCapacity(); err != nil {
			return fmt.Errorf("failed to decrease capacity: %w", err)
		}
	}
	return nil
}

// Size returns the number of elements in the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (al *ArrayList[T]) Size() int64 {
	return al.listSize
}

// IsEmpty returns true if the list contains no elements.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (al *ArrayList[T]) IsEmpty() bool {
	return al.listSize == 0
}

// Clear removes all elements from the list.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (al *ArrayList[T]) Clear() {
	al.listCapacity = 0
	al.listSize = 0
	al.data = nil
}

// increaseCapacity doubles the capacity.
// Time Complexity: O(n)
// Space Complexity: O(n)
func (al *ArrayList[T]) increaseCapacity() error {
	newCapacity := al.listCapacity*2 + 1
	newData := make([]T, newCapacity)
	copy(newData, al.data[:al.listSize])
	al.data = newData
	al.listCapacity = newCapacity
	return nil
}

// decreaseCapacity reduces the capacity to save memory.
// Time Complexity: O(n)
// Space Complexity: O(n)
func (al *ArrayList[T]) decreaseCapacity() error {
	newCapacity := al.listCapacity / 2
	if newCapacity < al.listSize {
		newCapacity = al.listSize
	}
	newData := make([]T, newCapacity)
	copy(newData, al.data[:al.listSize])
	al.data = newData
	al.listCapacity = newCapacity
	return nil
}

// checkIndexRange validates that index is within [0, listSize).
func (al *ArrayList[T]) checkIndexRange(index int64) error {
	if index < 0 || index >= al.listSize {
		return fmt.Errorf("index %d out of range [0, %d)", index, al.listSize)
	}
	return nil
}
