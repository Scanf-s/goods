package list

// List is the common interface for all list datastructures.
// Generic L could be slice([]) or Node pointer
// Generic T could be various builtin types or struct
type List[T any] interface {

	// Append adds an element to the end of the list.
	Append(element T) error

	// AppendAll adds all elements to the end of the list.
	// AppendAll(1, 2, 3) is equivalent to Append(1).Append(2).Append(3)
	AppendAll(elements ...T) error

	// Add inserts an element at the specific index in the list
	// It pushes all elements after the specific index to the right
	Add(index int, element T) error

	// Set replaces the element at the specific index in the list with newElement
	Set(index int, newElement T) error

	// Get returns the element at the specific index in the list
	Get(index int) (T, error)

	// Delete removes the element at the specific index in the list
	Delete(index int) error

	// Size returns the number of elements in the list
	Size() int

	// IsEmpty returns true if the list is empty
	IsEmpty() bool

	// Clear removes all elements from the list
	Clear() error
}
