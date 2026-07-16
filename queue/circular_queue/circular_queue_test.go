package circular_queue

import "testing"

func TestCircularQueue_New(t *testing.T) {
	queue := NewCircularQueue[int]()
	if queue == nil {
		t.Errorf("New linked list queue has to be initialized")
	}
}

func TestCircularQueue_Offer_Poll_Integers(t *testing.T) {
	queue := NewCircularQueue[int]()
	if queue == nil {
		t.Errorf("New circular queue has to be initialized")
	}

	var err error
	if err = queue.Offer(1); err != nil {
		t.Error(err)
	}
	if err = queue.Offer(2); err != nil {
		t.Error(err)
	}
	if err = queue.Offer(3); err != nil {
		t.Error(err)
	}

	var val int
	if val, err = queue.Poll(); err != nil {
		t.Error(err)
	}
	if val != 1 {
		t.Errorf("Poll expected 1, got %d", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Size expected 2, got %d", queue.Size())
	}

	if val, err = queue.Peek(); err != nil {
		t.Error(err)
	}
	if val != 2 {
		t.Errorf("Poll expected 2, got %d", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Size expected 2, got %d", queue.Size())
	}
	if val, err = queue.Poll(); err != nil {
		t.Error(err)
	}

	if val, err = queue.Poll(); err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Errorf("Poll expected 3, got %d", val)
	}
	if queue.Size() != 0 {
		t.Errorf("Size expected 0, got %d", queue.Size())
	}

	if !queue.IsEmpty() {
		t.Errorf("circular queue should be empty")
	}
}

func TestCircularQueue_EmptyOperations(t *testing.T) {
	queue := NewCircularQueue[int]()

	if !queue.IsEmpty() {
		t.Error("a new circular queue should be empty")
	}
	if queue.Size() != 0 {
		t.Errorf("a new circular queue should have size 0, got %d", queue.Size())
	}

	if _, err := queue.Peek(); err == nil {
		t.Error("Peek on an empty circular queue should return an error")
	}
	if _, err := queue.Poll(); err == nil {
		t.Error("Poll on an empty circular queue should return an error")
	}
}

func TestCircularQueue_FIFO_Order(t *testing.T) {
	queue := NewCircularQueue[int]()

	for i := 1; i <= 5; i++ {
		if err := queue.Offer(i); err != nil {
			t.Errorf("Offer failed: %s", err)
		}
	}

	// A queue is FIFO: elements come out in the order they went in.
	for _, want := range []int{1, 2, 3, 4, 5} {
		got, err := queue.Poll()
		if err != nil {
			t.Fatalf("Poll failed: %s", err)
		}
		if got != want {
			t.Errorf("Poll = %d, expected %d", got, want)
		}
	}

	if !queue.IsEmpty() {
		t.Errorf("queue should be empty after draining, size = %d", queue.Size())
	}
}

func TestCircularQueue_SingleElement(t *testing.T) {
	queue := NewCircularQueue[int]()

	if err := queue.Offer(99); err != nil {
		t.Errorf("Offer failed: %s", err)
	}
	if queue.Size() != 1 {
		t.Errorf("Size = %d, expected 1", queue.Size())
	}

	if val, err := queue.Peek(); err != nil || val != 99 {
		t.Errorf("Peek = %d (err %v), expected 99", val, err)
	}

	val, err := queue.Poll()
	if err != nil {
		t.Errorf("Poll failed: %s", err)
	}
	if val != 99 {
		t.Errorf("Poll = %d, expected 99", val)
	}
	if !queue.IsEmpty() {
		t.Errorf("queue should be empty after polling the only element, size = %d", queue.Size())
	}

	// Once empty again, reads must error.
	if _, err := queue.Peek(); err == nil {
		t.Error("Peek on the now-empty queue should return an error")
	}
	if _, err := queue.Poll(); err == nil {
		t.Error("Poll on the now-empty queue should return an error")
	}
}

func TestCircularQueue_WithStrings(t *testing.T) {
	queue := NewCircularQueue[string]()

	for _, s := range []string{"a", "b", "c"} {
		if err := queue.Offer(s); err != nil {
			t.Errorf("Offer failed: %s", err)
		}
	}

	for _, want := range []string{"a", "b", "c"} {
		got, err := queue.Poll()
		if err != nil {
			t.Fatalf("Poll failed: %s", err)
		}
		if got != want {
			t.Errorf("Poll = %q, expected %q", got, want)
		}
	}
}
