package ring_buffer

import "testing"

func TestRingBuffer_New(t *testing.T) {
	buffer := NewRingBuffer[int](5)
	if buffer == nil {
		t.Errorf("New linked list buffer has to be initialized")
	}
}

func TestRingBuffer_Offer_Poll_Integers(t *testing.T) {
	buffer := NewRingBuffer[int](5)
	if buffer == nil {
		t.Errorf("New ring buffer has to be initialized")
	}

	var err error
	if err = buffer.Offer(1); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(2); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(3); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(4); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(5); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(6); err == nil {
		t.Error("expect max buffer error but nil returned")
	}

	var val int
	if val, err = buffer.Peek(); err != nil {
		t.Error(err)
	}
	if val != 1 {
		t.Errorf("Poll expected 1, got %d", val)
	}
	if buffer.Size() != 5 {
		t.Errorf("Size expected 5, got %d", buffer.Size())
	}

	if val, err = buffer.Poll(); err != nil {
		t.Error(err)
	}
	if val != 1 {
		t.Errorf("Poll expected 1, got %d", val)
	}
	if buffer.Size() != 4 {
		t.Errorf("Size expected 4, got %d", buffer.Size())
	}

	if _, err = buffer.Poll(); err != nil {
		t.Error(err)
	}
	if val, err = buffer.Poll(); err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Errorf("Poll expected 3, got %d", val)
	}

	if err = buffer.Offer(123); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(456); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(789); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(1010101); err == nil {
		t.Error("expect max buffer error but nil returned")
	}
	if err = buffer.Offer(5679567856); err == nil {
		t.Error("expect max buffer error but nil returned")
	}
	if err = buffer.Offer(24573456); err == nil {
		t.Error("expect max buffer error but nil returned")
	}
	if err = buffer.Offer(9999); err == nil {
		t.Error("expect max buffer error but nil returned")
	}

	for range(5){
		if _, err = buffer.Poll(); err != nil {
			t.Error(err)
		}
	}
	if !buffer.IsEmpty() {
		t.Error("ring buffer should be emptied")
	}
	if _, err = buffer.Peek(); err == nil {
		t.Error("Peek on an empty ring buffer should return an error")
	}
	if val, err = buffer.Poll(); err == nil {
		t.Error("Poll on an empty ring buffer should return an error")
	}

	if err = buffer.Offer(1); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(2); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(3); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(4); err != nil {
		t.Error(err)
	}
	if err = buffer.Offer(5); err != nil {
		t.Error(err)
	}
	if buffer.Size() != 5 {
		t.Errorf("expected size of buffer: 5 but got %d", buffer.Size())
	}
}

func TestRingBuffer_EmptyOperations(t *testing.T) {
	buffer := NewRingBuffer[int](5)

	if !buffer.IsEmpty() {
		t.Error("a new ring buffer should be empty")
	}
	if buffer.Size() != 0 {
		t.Errorf("a new ring buffer should have size 0, got %d", buffer.Size())
	}

	if _, err := buffer.Peek(); err == nil {
		t.Error("Peek on an empty ring buffer should return an error")
	}
	if _, err := buffer.Poll(); err == nil {
		t.Error("Poll on an empty ring buffer should return an error")
	}
}

func TestRingBuffer_FIFO_Order(t *testing.T) {
	buffer := NewRingBuffer[int](10)

	for i := 1; i <= 10; i++ {
		if err := buffer.Offer(i); err != nil {
			t.Errorf("Offer failed: %s", err)
		}
	}

	// A buffer is FIFO: elements come out in the order they went in.
	for _, want := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		got, err := buffer.Poll()
		if err != nil {
			t.Fatalf("Poll failed: %s", err)
		}
		if got != want {
			t.Errorf("Poll = %d, expected %d", got, want)
		}
	}

	if !buffer.IsEmpty() {
		t.Errorf("buffer should be empty after draining, size = %d", buffer.Size())
	}
}

func TestRingBuffer_SingleElement(t *testing.T) {
	buffer := NewRingBuffer[int](1)

	if err := buffer.Offer(99); err != nil {
		t.Errorf("Offer failed: %s", err)
	}
	if buffer.Size() != 1 {
		t.Errorf("Size = %d, expected 1", buffer.Size())
	}

	if val, err := buffer.Peek(); err != nil || val != 99 {
		t.Errorf("Peek = %d (err %v), expected 99", val, err)
	}

	val, err := buffer.Poll()
	if err != nil {
		t.Errorf("Poll failed: %s", err)
	}
	if val != 99 {
		t.Errorf("Poll = %d, expected 99", val)
	}
	if !buffer.IsEmpty() {
		t.Errorf("buffer should be empty after polling the only element, size = %d", buffer.Size())
	}

	// Once empty again, reads must error.
	if _, err := buffer.Peek(); err == nil {
		t.Error("Peek on the now-empty buffer should return an error")
	}
	if _, err := buffer.Poll(); err == nil {
		t.Error("Poll on the now-empty buffer should return an error")
	}
}

func TestRingBuffer_WithStrings(t *testing.T) {
	buffer := NewRingBuffer[string](3)

	for _, s := range []string{"a", "b", "c"} {
		if err := buffer.Offer(s); err != nil {
			t.Errorf("Offer failed: %s", err)
		}
	}

	for _, want := range []string{"a", "b", "c"} {
		got, err := buffer.Poll()
		if err != nil {
			t.Fatalf("Poll failed: %s", err)
		}
		if got != want {
			t.Errorf("Poll = %q, expected %q", got, want)
		}
	}
}

func BenchmarkRingBuffer(b *testing.B) {
	buf := NewRingBuffer[int](1024)
	b.ReportAllocs()

	i := 0
	for b.Loop() {
		buf.Offer(i)
		buf.Poll()
		i++
	}
}
