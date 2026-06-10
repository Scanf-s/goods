package queue

type (
	Queue[T any] interface {
		// Offer adds new element into the queue data structure
		Offer(element T) error

		// Peek returns an element in front of the queue without removing it.
		Peek() (T, error)

		// Poll is returns an element in front of the queue as well as removes the front element.
		Poll() (T, error)
	}
)
