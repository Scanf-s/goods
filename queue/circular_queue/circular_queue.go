package circular_queue

import (
	"fmt"

	"github.com/Scanf-s/goods/list/linkedlist"
	"github.com/Scanf-s/goods/queue"
)

type CircularQueue[T any] struct {
	list *linkedlist.CircularLinkedList[T]
}

// NewCircularQueue initializes the circular linked list based queue
func NewCircularQueue[T any]() *CircularQueue[T] {
	return &CircularQueue[T]{
		list: linkedlist.NewCircularLinkedList[T](),
	}
}

func (c CircularQueue[T]) Offer(element T) error {
	if c.list == nil {
		return fmt.Errorf("you didn't initialize the circular queue. please call NewCircularQueue() first")
	}
	err := c.list.Append(element)
	if err != nil {
		return fmt.Errorf("failed to append item into circular queue")
	}
	return nil
}

func (c CircularQueue[T]) Peek() (T, error) {
	var element T
	val, err := c.list.Head()
	if err != nil {
		return element, fmt.Errorf("failed to peek item from circular queue. The list is empty")
	}
	return val, nil
}

func (c CircularQueue[T]) Poll() (T, error) {
	var element T
	val, err := c.list.PopHead()
	if err != nil {
		return element, fmt.Errorf("failed to poll item from circular queue. The list is empty")
	}
	return val, nil
}

func (c CircularQueue[T]) Size() int {
	return c.list.Size()
}

func (c CircularQueue[T]) IsEmpty() bool {
	return c.list.IsEmpty()
}

var _ queue.Queue[int] = (*CircularQueue[int])(nil)
