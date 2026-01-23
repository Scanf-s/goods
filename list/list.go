package list

// List is the common interface for all list datastructures.
// Generic L could be slice([]) or Node pointer
// Generic T could be various builtin types or struct
type List[T any] interface {

	// Append adds an element to the end of the list.
	Append(element T) error

	// Add inserts an element at the specific index in the list
	Add(index int, element T) error

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
