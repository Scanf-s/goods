package ring_buffer

import (
	"fmt"

	"github.com/Scanf-s/goods/queue"
)

type RingBuffer[T any] struct {

	// buffer represents data storage
	buffer []T

	// front represents the head of the queue
	front int

	// rear represents the rear of the queue
	rear int

	// count represents existing a number of elements in buffer
	count int
}

// NewRingBuffer initializes the fixed-lenght slice based queue
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		panic("capacity must be positive")	
	}
	return &RingBuffer[T]{
		buffer: make([]T, capacity),
		front: 0,
		rear: 0,
		count: 0,
	}
}

func (r *RingBuffer[T]) Offer(element T) error {
	if r.buffer == nil {
		return fmt.Errorf("you didn't initialize the ring buffer. please call NewRingBuffer() first")
	}

	// check buffer size
	if r.count == cap(r.buffer) {
		return fmt.Errorf("buffer is full")
	}

	// store element
	r.buffer[r.rear] = element
	r.rear = (r.rear + 1) % cap(r.buffer)
	r.count++
	return nil
}

func (r *RingBuffer[T]) Peek() (T, error) {
	var element T
	if r.IsEmpty() {
		return element, fmt.Errorf("buffer is empty")
	}
	return r.buffer[r.front], nil
}

func (r *RingBuffer[T]) Poll() (T, error) {
	var zero T
	if r.IsEmpty() {
		return zero, fmt.Errorf("buffer is empty")
	}
	val := r.buffer[r.front]

	// set default value on emptied location
	r.buffer[r.front] = zero

	r.front = (r.front + 1) % cap(r.buffer)
	r.count--
	return val, nil
}

func (r *RingBuffer[T]) Size() int {
	return r.count
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.count == 0
}

var _ queue.Queue[int] = (*RingBuffer[int])(nil)
