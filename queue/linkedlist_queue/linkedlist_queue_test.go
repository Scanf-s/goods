package linkedlist_queue

import "testing"

func TestLinkedListQueue_New(t *testing.T) {
	queue := NewLinkedListQueue[int]()
	if queue == nil {
		t.Errorf("New stack array has to be initialized")
	}
}

func TestLinkedListQueue_Push_Integers(t *testing.T) {
	queue := NewLinkedListQueue[int]()
	if queue == nil {
		t.Errorf("New stack array has to be initialized")
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

	if val, err = queue.Poll(); err != nil {
		t.Error(err)
	}
	if val != 2 {
		t.Errorf("Poll expected 2, got %d", val)
	}

	if val, err = queue.Poll(); err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Errorf("Poll expected 3, got %d", val)
	}
}
