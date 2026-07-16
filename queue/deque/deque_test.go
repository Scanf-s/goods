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

func TestDeque_EmptyOperations(t *testing.T) {
	deque := NewDeque[int]()

	if !deque.IsEmpty() {
		t.Error("a new deque should be empty")
	}
	if deque.Size() != 0 {
		t.Errorf("a new deque should have size 0, got %d", deque.Size())
	}

	// Every read/remove on an empty deque must error, not panic.
	if _, err := deque.Peek(); err == nil {
		t.Error("Peek on an empty deque should return an error")
	}
	if _, err := deque.PeekFront(); err == nil {
		t.Error("PeekFront on an empty deque should return an error")
	}
	if _, err := deque.Poll(); err == nil {
		t.Error("Poll on an empty deque should return an error")
	}
	if _, err := deque.PollFront(); err == nil {
		t.Error("PollFront on an empty deque should return an error")
	}
}

func TestDeque_OfferFront_Ordering(t *testing.T) {
	deque := NewDeque[int]()

	// OfferFront pushes to the front, Offer pushes to the back.
	// After these calls the logical order (front -> back) is [3, 2, 1, 4].
	if err := deque.OfferFront(1); err != nil {
		t.Error(err)
	}
	if err := deque.OfferFront(2); err != nil {
		t.Error(err)
	}
	if err := deque.OfferFront(3); err != nil {
		t.Error(err)
	}
	if err := deque.Offer(4); err != nil {
		t.Error(err)
	}

	if deque.Size() != 4 {
		t.Errorf("Size = %d, expected 4", deque.Size())
	}

	front, err := deque.PeekFront()
	if err != nil {
		t.Error(err)
	}
	if front != 3 {
		t.Errorf("PeekFront = %d, expected 3", front)
	}

	back, err := deque.Peek()
	if err != nil {
		t.Error(err)
	}
	if back != 4 {
		t.Errorf("Peek (back) = %d, expected 4", back)
	}

	// Remove from the front twice: 3 then 2.
	if v, err := deque.PollFront(); err != nil || v != 3 {
		t.Errorf("PollFront = %d (err %v), expected 3", v, err)
	}
	if v, err := deque.PollFront(); err != nil || v != 2 {
		t.Errorf("PollFront = %d (err %v), expected 2", v, err)
	}

	// Remove from the back: 4.
	if v, err := deque.Poll(); err != nil || v != 4 {
		t.Errorf("Poll (back) = %d (err %v), expected 4", v, err)
	}

	// Only element 1 should remain.
	if deque.Size() != 1 {
		t.Errorf("Size = %d, expected 1", deque.Size())
	}
}

func TestDeque_SingleElement_PollFront(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PollFront on a single-element deque should not panic: %v", r)
		}
	}()

	deque := NewDeque[int]()
	if err := deque.Offer(7); err != nil {
		t.Error(err)
	}

	val, err := deque.PollFront()
	if err != nil {
		t.Error(err)
	}
	if val != 7 {
		t.Errorf("PollFront = %d, expected 7", val)
	}
	if !deque.IsEmpty() {
		t.Errorf("deque should be empty after polling the only element, size = %d", deque.Size())
	}
}

func TestDeque_SingleElement_Poll(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Poll on a single-element deque should not panic: %v", r)
		}
	}()

	deque := NewDeque[int]()
	if err := deque.OfferFront(7); err != nil {
		t.Error(err)
	}

	val, err := deque.Poll()
	if err != nil {
		t.Error(err)
	}
	if val != 7 {
		t.Errorf("Poll = %d, expected 7", val)
	}
	if !deque.IsEmpty() {
		t.Errorf("deque should be empty after polling the only element, size = %d", deque.Size())
	}
}

func TestDeque_WithStrings(t *testing.T) {
	deque := NewDeque[string]()

	if err := deque.Offer("b"); err != nil {
		t.Error(err)
	}
	if err := deque.OfferFront("a"); err != nil {
		t.Error(err)
	}
	if err := deque.Offer("c"); err != nil {
		t.Error(err)
	}

	// Logical order (front -> back): [a, b, c].
	if front, err := deque.PeekFront(); err != nil || front != "a" {
		t.Errorf("PeekFront = %q (err %v), expected \"a\"", front, err)
	}
	if back, err := deque.Peek(); err != nil || back != "c" {
		t.Errorf("Peek (back) = %q (err %v), expected \"c\"", back, err)
	}
}
