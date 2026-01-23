package arraylist

import (
	"fmt"
)

// ArrayList represents the ArrayList datastructure
type ArrayList[T any] struct {
	// data is the array to store the T type elements
	data []T

	// listCapacity represents the capacity of the list (total number of elements that could be stored)
	listCapacity int64

	// listSize represents the number of elements currently stored in the list
	listSize int64
}

// New returns a new specific sized ArrayList
func New[T any](size int64) *ArrayList[T] {
	return &ArrayList[T]{
		data:         make([]T, size),
		listCapacity: size,
		listSize:     size,
	}
}

func (al *ArrayList[T]) Append(element T) error {
	if ok := al.checkAvailableSpace(element); !ok {
		// Increase capacity if there is not enough space
		err := al.increaseCapacity()
		if err != nil {
			return fmt.Errorf("failed to increase capacity: %w", err)
		}
	}
	al.data = append(al.data, element)
	return nil
}

func (al *ArrayList[T]) AppendAll(elements ...T) error {
	if ok := al.checkAvailableSpace(elements...); !ok {
		// Increase capacity if there is not enough space
		err := al.increaseCapacity()
		if err != nil {
			return fmt.Errorf("failed to increase capacity: %w", err)
		}
	}
	al.data = append(al.data, elements...)
	al.listSize += int64(len(elements))
	return nil
}

func (al *ArrayList[T]) Add(index int64, element T) error {
	if err := al.checkIndexRange(index); err != nil {
		return err
	}

	// Insert an element at the specific index
	al.data[index] = element
	return nil
}

func (al *ArrayList[T]) Get(index int64) (any, error) {
	if err := al.checkIndexRange(index); err != nil {
		return nil, err
	}

	return al.data[index], nil
}

func (al *ArrayList[T]) Delete(index int64) error {
	if err := al.checkIndexRange(index); err != nil {
		return err
	}

	// Delete an element at the specific index and arrange the list
	copy(al.data[index:], al.data[index+1:])
	al.data = al.data[:al.listSize-1]
	al.listSize--

	// Check if the current size of the list is lower than the half of the capacity
	if (al.listSize * 2) <= al.listCapacity {
		// Reduce the capacity of the list if it is lower than half of the capacity
		err := al.decreaseCapacity()
		if err != nil {
			return fmt.Errorf("failed to decrease capacity: %w", err)
		}
	}
	return nil
}

func (al *ArrayList[T]) Size() int64 {
	return al.listSize
}

func (al *ArrayList[T]) IsEmpty() bool {
	return al.listSize == 0
}

func (al *ArrayList[T]) Clear() error {
	al.listCapacity = 0
	al.listSize = 0
	al.data = nil
	return nil
}

// increaseCapacity increases the capacity of the array automatically
// It increases the capacity by CUR_ARRAY_LIST_SIZE * 2 + 1
func (al *ArrayList[T]) increaseCapacity() error {
	// Create a new dynamic array
	newCapacity := al.listCapacity*2 + 1
	newData := make([]T, newCapacity)

	// use a builtin copy function to move all elements from an old array to a new array
	copy(newData, al.data)
	al.data = newData
	al.listCapacity = newCapacity
	return nil
}

func (al *ArrayList[T]) decreaseCapacity() error {
	// Create a new dynamic array
	newCapacity := int64(al.listCapacity / 2)
	newData := make([]T, newCapacity)

	// copy elements into new array
	copy(newData, al.data[:al.listSize])
	al.data = newData
	al.listCapacity = newCapacity
	return nil
}

func (al *ArrayList[T]) checkIndexRange(index int64) error {
	// Check index range
	if index >= al.listCapacity || index < 0 {
		return fmt.Errorf("index %d out of range [0, %d]", index, al.listCapacity)
	}
	return nil
}

// checkAvailableSpace checks if there is enough space to insert the elements
func (al *ArrayList[T]) checkAvailableSpace(elements ...T) bool {
	// Check available space before insert the elements
	availableSpaces := cap(al.data) - len(al.data)
	if availableSpaces < len(elements) {
		return false
	}
	return true
}
