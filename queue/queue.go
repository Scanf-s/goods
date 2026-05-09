package queue

type (
	Queue[T any] interface {
		// Enqueue is ...
		Enqueue(element T) error

		// Peek is ...
		Peek() (T, error)

		// FrontPeek is ...
		FrontPeek() (T, error)

		// Dequeue is ...
		Dequeue() (T, error)
	}
)
