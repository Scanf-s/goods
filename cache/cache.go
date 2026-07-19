package cache

type (

	// Node of Cached data
	// This node will support O(1) Get, Put operations.
	Node[K comparable, V any] struct {
		Key  K
		Data V
		Prev *Node[K, V]
		Next *Node[K, V]
	}

	// Cache interface
	Cache[K comparable, V any] interface {
		// Get returns cached value from a storage
		Get(key K) (V, bool)

		// Put will store data using key into a storage
		Put(key K, data V)

		UpdateCapacity(capacity int) error
	}
)
