package arraylistqueue

import (
	"fmt"
	"github.com/Scanf-s/goods/list/arraylist"
	"github.com/Scanf-s/goods/queue"
)

// ArrayListQueue is a FIFO queue built by composing an ArrayList.
//
// Because it reuses ArrayList (a positional list anchored at index 0), removing
// the front element means deleting index 0, which shifts every remaining element
// left. That makes Poll O(n): the queue inherits the cost model of its substrate.
// For an O(1) array-backed queue, see CircularQueue, which uses a ring buffer with
// head/tail indices over a raw slice instead of a list.
type ArrayListQueue[T any] struct {
	list *arraylist.ArrayList[T]
}

// Compile-time check that ArrayListQueue implements the queue.Queue interface.
var _ queue.Queue[int] = (*ArrayListQueue[int])(nil)

// NewArrayListQueue initializes the queue that uses array(dynamic) list for its base data storage.
// The size argument is an initial capacity hint forwarded to the backing ArrayList.
func NewArrayListQueue[T any](size int) *ArrayListQueue[T] {
	return &ArrayListQueue[T]{
		list: arraylist.New[T](size),
	}
}

// Offer adds an element to the back of the queue.
// Time Complexity: O(1) amortized (O(n) when the backing array resizes)
func (alq *ArrayListQueue[T]) Offer(value T) error {
	if alq == nil {
		return fmt.Errorf("queue is nil")
	}
	if err := alq.list.Append(value); err != nil {
		return fmt.Errorf("failed to offer an element into the queue")
	}
	return nil
}

// Peek returns the front element without removing it.
// Returns an error if the queue is empty.
// Time Complexity: O(1)
func (alq *ArrayListQueue[T]) Peek() (T, error) {
	var defaultVal T
	if alq == nil {
		return defaultVal, fmt.Errorf("queue is nil")
	}
	value, err := alq.list.Get(0)
	if err != nil {
		return defaultVal, fmt.Errorf("cannot peek front element from the queue")
	}
	return value, nil
}

// Poll returns and removes the front element.
// Returns an error if the queue is empty.
//
// Time Complexity: O(n). Removing the front delegates to ArrayList.Delete(0),
// which shifts every remaining element one slot left. This O(n) cost is a
// deliberate consequence of reusing ArrayList.
func (alq *ArrayListQueue[T]) Poll() (T, error) {
	var defaultVal T
	if alq == nil {
		return defaultVal, fmt.Errorf("queue is nil")
	}

	value, err := alq.list.Get(0)
	if err != nil {
		return defaultVal, fmt.Errorf("cannot poll front element from the queue")
	}
	if err := alq.list.Delete(0); err != nil {
		return defaultVal, fmt.Errorf("failed to remove front element from the queue")
	}
	return value, nil
}

// Size returns the number of elements in the queue.
// Time Complexity: O(1)
func (alq *ArrayListQueue[T]) Size() int {
	if alq == nil {
		return 0
	}
	return alq.list.Size()
}

// IsEmpty returns true if the queue has no elements.
// Time Complexity: O(1)
func (alq *ArrayListQueue[T]) IsEmpty() bool {
	if alq == nil {
		return true
	}
	return alq.list.IsEmpty()
}
