package deque

import (
	"fmt"

	"github.com/Scanf-s/goods/list/linkedlist"
	"github.com/Scanf-s/goods/queue"
)

type Deque[T any] struct {
	list *linkedlist.DoublyLinkedList[T]
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		list: linkedlist.NewDoublyLinkedList[T](),
	}
}

func (d Deque[T]) Offer(element T) error {
	if d.list == nil {
		return fmt.Errorf("you didn't initialize the list. please call NewDeque() first")
	}
	err := d.list.Append(element)
	if err != nil {
		return fmt.Errorf("there was an error adding the element to the list")
	}
	return nil
}

func (d Deque[T]) OfferFront(element T) error {
	if d.list == nil {
		return fmt.Errorf("you didn't initialize the list. please call NewDeque() first")
	}
	err := d.list.Prepend(element)
	if err != nil {
		return fmt.Errorf("there was an error adding the element to the list")
	}
	return nil
}

func (d Deque[T]) Peek() (T, error) {
	var element T
	if d.list == nil {
		return element, fmt.Errorf("deque is nil")
	}

	element, err := d.list.Tail()
	if err != nil {
		return element, fmt.Errorf("cannot peek front element from the queue")
	}
	return element, nil
}

func (d Deque[T]) Poll() (T, error) {
	var element T
	if d.list == nil {
		return element, fmt.Errorf("deque is nil")
	}

	element, err := d.list.PopTail()
	if err != nil {
		return element, fmt.Errorf("cannot peek front element from the queue")
	}
	return element, nil
}

func (d Deque[T]) PeekFront() (T, error) {
	var element T
	if d.list == nil {
		return element, fmt.Errorf("deque is nil")
	}

	element, err := d.list.Head()
	if err != nil {
		return element, fmt.Errorf("cannot peek front element from the queue")
	}
	return element, nil
}

func (d Deque[T]) PollFront() (T, error) {
	var element T
	if d.list == nil {
		return element, fmt.Errorf("deque is nil")
	}

	element, err := d.list.PopHead()
	if err != nil {
		return element, fmt.Errorf("cannot peek front element from the queue")
	}
	return element, nil
}

func (d Deque[T]) Size() int {
	return d.list.Size()
}

func (d Deque[T]) IsEmpty() bool {
	return d.list.IsEmpty()
}

var _ queue.Queue[int] = (*Deque[int])(nil)
