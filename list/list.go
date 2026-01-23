package list

type (
	// Node represents a node in a linked list.
	Node struct {
		data any
		next *Node
	}

	// List is a linked list data structure.
	List struct {
		head *Node
		tail *Node
		size int
	}

	// ListOperation is an interface for a list data structure.
	ListOperation interface {
		// Append function add elements in the tail of the List
		Append(list *List, element any)

		// Add function add elements in the specific location of the List
		Add(list *List, index int, element any)

		// Delete function delete element in the specific location of the List
		Delete(list *List, index int)

		// Get retrieves the node at the specified index.
		Get(list *List, index int)

		// IsEmpty checks if the List is empty.
		IsEmpty(list *List)
	}
)