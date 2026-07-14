package deque

import "testing"

func TestDeque_New(t *testing.T) {
	deque := NewDeque[int]()
	if deque == nil {
		t.Errorf("New deque has to be initialized")
	}
}

func TestDeque_Offer_Poll_Integers(t *testing.T) {
	deque := NewDeque[int]()
	if deque == nil {
		t.Errorf("New deque has to be initialized")
	}

	var err error
	if err = deque.Offer(1); err != nil {
		t.Error(err)
	}
	if err = deque.Offer(2); err != nil {
		t.Error(err)
	}
	if err = deque.Offer(3); err != nil {
		t.Error(err)
	}

	var val int
	if val, err = deque.Poll(); err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Errorf("Poll expected 3, got %d", val)
	}

	if val, err = deque.PollFront(); err != nil {
		t.Error(err)
	}
	if val != 1 {
		t.Errorf("PollFront expected 1, got %d", val)
	}

	if val, err = deque.Peek(); err != nil {
		t.Error(err)
	}
	if val != 2 {
		t.Errorf("Poll expected 2, got %d", val)
	}
}
