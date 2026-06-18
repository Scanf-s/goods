package arraylistqueue

import "testing"

func TestArrayListQueue_New(t *testing.T) {
	queue := NewArrayListQueue[int](0)
	if queue == nil {
		t.Fatal("New array list queue has to be initialized")
	}
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if queue.Size() != 0 {
		t.Errorf("New queue size = %d, expected 0", queue.Size())
	}
}

func TestArrayListQueue_Offer_Poll_FIFO(t *testing.T) {
	queue := NewArrayListQueue[int](0)

	for _, v := range []int{1, 2, 3} {
		if err := queue.Offer(v); err != nil {
			t.Errorf("Offer(%d) failed: %s", v, err)
		}
	}
	if queue.Size() != 3 {
		t.Errorf("Size = %d, expected 3", queue.Size())
	}

	// FIFO: elements come out in the order they went in.
	for _, want := range []int{1, 2, 3} {
		val, err := queue.Poll()
		if err != nil {
			t.Errorf("Poll failed: %s", err)
		}
		if val != want {
			t.Errorf("Poll = %d, expected %d", val, want)
		}
	}
	if !queue.IsEmpty() {
		t.Error("queue should be empty after polling all elements")
	}
}

func TestArrayListQueue_Peek(t *testing.T) {
	queue := NewArrayListQueue[int](0)
	if err := queue.Offer(10); err != nil {
		t.Errorf("Offer failed: %s", err)
	}
	if err := queue.Offer(20); err != nil {
		t.Errorf("Offer failed: %s", err)
	}

	// Peek returns the front without removing it.
	val, err := queue.Peek()
	if err != nil {
		t.Errorf("Peek failed: %s", err)
	}
	if val != 10 {
		t.Errorf("Peek = %d, expected 10", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Size changed after Peek: %d, expected 2", queue.Size())
	}
}

func TestArrayListQueue_Peek_Empty(t *testing.T) {
	queue := NewArrayListQueue[int](0)
	if _, err := queue.Peek(); err == nil {
		t.Error("Peek on empty queue should return an error")
	}
}

func TestArrayListQueue_Poll_Empty(t *testing.T) {
	queue := NewArrayListQueue[int](0)
	if _, err := queue.Poll(); err == nil {
		t.Error("Poll on empty queue should return an error")
	}
}

func TestArrayListQueue_InterleavedOps(t *testing.T) {
	queue := NewArrayListQueue[int](0)

	_ = queue.Offer(1)
	_ = queue.Offer(2)
	if val, _ := queue.Poll(); val != 1 { // remove 1 -> [2]
		t.Errorf("Poll = %d, expected 1", val)
	}
	_ = queue.Offer(3) // [2, 3]
	_ = queue.Offer(4) // [2, 3, 4]
	if val, _ := queue.Poll(); val != 2 {
		t.Errorf("Poll = %d, expected 2", val)
	}
	if val, _ := queue.Peek(); val != 3 {
		t.Errorf("Peek = %d, expected 3", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Size = %d, expected 2", queue.Size())
	}
}

func TestArrayListQueue_WithStrings(t *testing.T) {
	queue := NewArrayListQueue[string](0)
	_ = queue.Offer("a")
	_ = queue.Offer("b")

	if val, _ := queue.Poll(); val != "a" {
		t.Errorf("Poll = %q, expected \"a\"", val)
	}
	if val, _ := queue.Poll(); val != "b" {
		t.Errorf("Poll = %q, expected \"b\"", val)
	}
}
