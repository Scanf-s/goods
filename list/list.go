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
	Add(index int64, element T) error

	// Get returns the element at the specific index in the list
	Get(index int64) (T, error)

	// Delete removes the element at the specific index in the list
	Delete(index int64) error

	// Size returns the number of elements in the list
	Size() int64

	// IsEmpty returns true if the list is empty
	IsEmpty() bool

	// Clear removes all elements from the list
	Clear() error
}
