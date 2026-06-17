package linkedlist_queue

import (
	"fmt"
	"github.com/Scanf-s/goods/list/linkedlist"
	"github.com/Scanf-s/goods/queue"
)

type LinkedListQueue[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

var _ queue.Queue[int] = (*LinkedListQueue[int])(nil)

// NewLinkedListQueue initializes the queue that uses linked list for its base data storage.
func NewLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{
		list: linkedlist.NewSinglyLinkedList[T](),
	}
}

func (llq *LinkedListQueue[T]) Offer(value T) error {
	if llq == nil {
		return fmt.Errorf("queue is nil")
	}
	err := llq.list.Append(value)
	if err != nil {
		return fmt.Errorf("failed to append new value into the queue")
	}
	return nil
}

func (llq *LinkedListQueue[T]) Peek() (T, error) {
	var element T
	if llq == nil {
		return element, fmt.Errorf("queue is nil")
	}

	element, err := llq.list.Head()
	if err != nil {
		return element, fmt.Errorf("cannot peek front element from the queue")
	}
	return element, nil
}

func (llq *LinkedListQueue[T]) Poll() (T, error) {
	var element T
	if llq == nil {
		return element, fmt.Errorf("queue is nil")
	}

	element, err := llq.list.PopHead()
	if err != nil {
		return element, fmt.Errorf("cannot poll front element from the queue")
	}
	return element, nil
}

func (llq *LinkedListQueue[T]) Size() int {
	if llq == nil {
		return 0
	}
	return llq.list.Size()
}

func (llq *LinkedListQueue[T]) IsEmpty() bool {
	if llq == nil {
		return true
	}
	return llq.list.IsEmpty()
}
