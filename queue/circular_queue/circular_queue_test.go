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
