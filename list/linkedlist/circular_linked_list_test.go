package linkedlist

import "testing"

func TestCircularLinkedList_New(t *testing.T) {
	list := NewCircularLinkedList[int]()

	if list.Size() != 0 {
		t.Errorf("New list should have size 0, got %d", list.Size())
	}
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
	if list.head != nil {
		t.Error("New list head should be nil")
	}
	if list.tail != nil {
		t.Error("New list tail should be nil")
	}
}

func TestCircularLinkedList_Append(t *testing.T) {
	testCases := []struct {
		name         string
		appendCount  int
		expectedSize int
	}{
		{
			name:         "Append single element",
			appendCount:  1,
			expectedSize: 1,
		},
		{
			name:         "Append 10 elements",
			appendCount:  10,
			expectedSize: 10,
		},
		{
			name:         "Append 100 elements",
			appendCount:  100,
			expectedSize: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewCircularLinkedList[int]()

			for i := range tc.appendCount {
				err := list.Append(i)
				if err != nil {
					t.Errorf("Append failed: %s", err)
				}
			}

			if list.Size() != tc.expectedSize {
				t.Errorf("Expected size %d, got %d", tc.expectedSize, list.Size())
			}

			// Verify elements are stored correctly
			for i := range tc.appendCount {
				val, err := list.Get(i)
				if err != nil {
					t.Errorf("Get(%d) failed: %s", i, err)
				}
				if val != i {
					t.Errorf("Get(%d) = %d, expected %d", i, val, i)
				}
			}
		})
	}
}

func TestCircularLinkedList_Append_HeadTail(t *testing.T) {
	list := NewCircularLinkedList[int]()

	// First append should set both head and tail
	if err := list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if list.head != list.tail {
		t.Error("After first append, head and tail should be the same")
	}
	if list.head.Data != 1 {
		t.Errorf("Head data = %d, expected 1", list.head.Data)
	}

	// Second append should update tail only
	if err := list.Append(2); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if list.head == list.tail {
		t.Error("After second append, head and tail should be different")
	}
	if list.head.Data != 1 {
		t.Errorf("Head data = %d, expected 1", list.head.Data)
	}
	if list.tail.Data != 2 {
		t.Errorf("Tail data = %d, expected 2", list.tail.Data)
	}
}

func TestCircularLinkedList_Append_CircularProperty(t *testing.T) {
	list := NewCircularLinkedList[int]()

	// Single element should point to itself
	if err := list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if list.head.Next != list.head {
		t.Error("Single element's Next should point to itself")
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to head")
	}

	// Multiple elements: tail.Next should point to head
	if err := list.Append(2); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err := list.Append(3); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to head")
	}

	// Traverse full circle
	current := list.head
	for range list.Size() {
		current = current.Next
	}
	if current != list.head {
		t.Error("After traversing Size() times, should return to head")
	}
}

func TestCircularLinkedList_AppendAll(t *testing.T) {
	list := NewCircularLinkedList[int]()

	err := list.AppendAll(1, 2, 3, 4, 5)
	if err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	if list.Size() != 5 {
		t.Errorf("Expected size 5, got %d", list.Size())
	}

	// Append more
	err = list.AppendAll(6, 7, 8)
	if err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	if list.Size() != 8 {
		t.Errorf("Expected size 8, got %d", list.Size())
	}

	// Verify all elements
	for i := 1; i <= 8; i++ {
		val, _ := list.Get(i - 1)
		if val != i {
			t.Errorf("Get(%d) = %d, expected %d", i-1, val, i)
		}
	}

	// Verify circular property after AppendAll
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to head after AppendAll")
	}
}

func TestCircularLinkedList_Add(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append(3); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append(4); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	// list: [1, 3, 4]

	// Insert 2 at index 1
	err = list.Add(1, 2)
	if err != nil {
		t.Errorf("Add failed: %s", err)
	}
	// list: [1, 2, 3, 4]

	if list.Size() != 4 {
		t.Errorf("Expected size 4, got %d", list.Size())
	}

	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Get(%d) = %d, expected %d", i, val, exp)
		}
	}

	// Verify circular property
	if list.tail.Next != list.head {
		t.Error("Circular property broken after Add")
	}
}

func TestCircularLinkedList_Add_AtHead(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(2, 3, 4); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Insert at the beginning
	err = list.Add(0, 1)
	if err != nil {
		t.Errorf("Add at beginning failed: %s", err)
	}

	if list.head.Data != 1 {
		t.Errorf("Head should be updated to 1, got %d", list.head.Data)
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to new head")
	}

	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Get(%d) = %d, expected %d", i, val, exp)
		}
	}
}

func TestCircularLinkedList_Add_EmptyList(t *testing.T) {
	list := NewCircularLinkedList[int]()

	// Add(0) into an empty list seeds it (consistent with Singly/Doubly).
	if err := list.Add(0, 1); err != nil {
		t.Errorf("Add(0,1) on empty list should succeed, got: %s", err)
	}
	if list.Size() != 1 {
		t.Errorf("Size = %d, expected 1", list.Size())
	}
	if val, _ := list.Get(0); val != 1 {
		t.Errorf("Get(0) = %d, expected 1", val)
	}
	// Single element must still form a closed ring.
	if list.head.Next != list.head {
		t.Error("Single element's Next should point to itself")
	}
}

func TestCircularLinkedList_Add_Append(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if err := list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Add at index == Size() appends to the end.
	if err := list.Add(3, 4); err != nil {
		t.Errorf("Add(3,4) should append, got: %s", err)
	}
	if list.tail.Data != 4 {
		t.Errorf("Tail = %d, expected 4", list.tail.Data)
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to head after append via Add")
	}

	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		if val, _ := list.Get(i); val != exp {
			t.Errorf("Get(%d) = %d, expected %d", i, val, exp)
		}
	}
}

func TestCircularLinkedList_Set(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	err = list.Set(1, 20)
	if err != nil {
		t.Errorf("Set failed: %s", err)
	}

	val, _ := list.Get(1)
	if val != 20 {
		t.Errorf("Get(1) = %d, expected 20", val)
	}

	// Size should not change
	if list.Size() != 3 {
		t.Errorf("Size changed after Set: expected 3, got %d", list.Size())
	}
}

func TestCircularLinkedList_Set_EmptyList(t *testing.T) {
	list := NewCircularLinkedList[int]()

	err := list.Set(0, 1)
	if err == nil {
		t.Error("Set on empty list should return error")
	}
}

func TestCircularLinkedList_Get_EmptyList(t *testing.T) {
	list := NewCircularLinkedList[int]()

	_, err := list.Get(0)
	if err == nil {
		t.Error("Get on empty list should return error")
	}
}

func TestCircularLinkedList_Get_OutOfRange(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if err := list.AppendAll(0, 1, 2); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}
	// list: [0, 1, 2]; valid indices are 0, 1, 2.

	// In-range indices succeed.
	for i := range 3 {
		val, err := list.Get(i)
		if err != nil {
			t.Errorf("Get(%d) failed: %s", i, err)
		}
		if val != i {
			t.Errorf("Get(%d) = %d, expected %d", i, val, i)
		}
	}

	// Out-of-range indices return an error (no silent wrapping).
	for _, idx := range []int{3, 4, 6, -1, -4} {
		if _, err := list.Get(idx); err == nil {
			t.Errorf("Get(%d) should return an out-of-range error", idx)
		}
	}
}

func TestCircularLinkedList_Delete(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3, 4, 5); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}
	// list: [1, 2, 3, 4, 5]

	// Delete middle element
	err = list.Delete(2)
	if err != nil {
		t.Errorf("Delete failed: %s", err)
	}
	// list: [1, 2, 4, 5]

	if list.Size() != 4 {
		t.Errorf("Expected size 4, got %d", list.Size())
	}

	expected := []int{1, 2, 4, 5}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Get(%d) = %d, expected %d", i, val, exp)
		}
	}

	// Verify circular property
	if list.tail.Next != list.head {
		t.Error("Circular property broken after Delete")
	}
}

func TestCircularLinkedList_Delete_Head(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Delete first element
	if err = list.Delete(0); err != nil {
		t.Errorf("Delete failed: %s", err)
	}

	if list.head.Data != 2 {
		t.Errorf("Head should be updated to 2, got %d", list.head.Data)
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to new head")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestCircularLinkedList_Delete_Tail(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Delete last element
	if err = list.Delete(list.Size() - 1); err != nil {
		t.Errorf("Delete failed: %s", err)
	}

	if list.tail.Data != 2 {
		t.Errorf("Tail should be updated to 2, got %d", list.tail.Data)
	}
	if list.tail.Next != list.head {
		t.Error("Tail's Next should point to head")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestCircularLinkedList_Delete_SingleElement(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	err = list.Delete(0)
	if err != nil {
		t.Errorf("Delete failed: %s", err)
	}

	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
	if list.head != nil {
		t.Error("Head should be nil after deleting single element")
	}
	if list.tail != nil {
		t.Error("Tail should be nil after deleting single element")
	}
}

func TestCircularLinkedList_Delete_EmptyList(t *testing.T) {
	list := NewCircularLinkedList[int]()

	err := list.Delete(0)
	if err == nil {
		t.Error("Delete on empty list should return error")
	}
}

func TestCircularLinkedList_Delete_OutOfRange(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if err := list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	if err := list.Delete(3); err == nil {
		t.Error("Delete(3) on a 3-element list should return an out-of-range error")
	}
	if err := list.Delete(-1); err == nil {
		t.Error("Delete(-1) should return an out-of-range error")
	}
}

func TestCircularLinkedList_Clear(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3, 4, 5); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	list.Clear()

	if list.Size() != 0 {
		t.Errorf("Size after Clear should be 0, got %d", list.Size())
	}
	if !list.IsEmpty() {
		t.Error("List should be empty after Clear")
	}
	if list.head != nil {
		t.Error("Head should be nil after Clear")
	}
	if list.tail != nil {
		t.Error("Tail should be nil after Clear")
	}
}

func TestCircularLinkedList_CircularTraversal(t *testing.T) {
	list := NewCircularLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Traverse 2 full circles (6 elements for size 3)
	expected := []int{1, 2, 3, 1, 2, 3}
	current := list.head
	for i, exp := range expected {
		if current.Data != exp {
			t.Errorf("Traversal[%d] = %d, expected %d", i, current.Data, exp)
		}
		current = current.Next
	}
}

func TestCircularLinkedList_WithStrings(t *testing.T) {
	list := NewCircularLinkedList[string]()

	var err error
	if err = list.Append("hello"); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append("world"); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Get failed: %s", err)
	}
	if val != "hello" {
		t.Errorf("Get(0) = %s, expected 'hello'", val)
	}

	val, _ = list.Get(1)
	if val != "world" {
		t.Errorf("Get(1) = %s, expected 'world'", val)
	}

	// Out-of-range index returns an error (no wrapping).
	if _, err = list.Get(2); err == nil {
		t.Error("Get(2) on a 2-element list should return an out-of-range error")
	}
}

func TestCircularLinkedList_WithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	var err error
	list := NewCircularLinkedList[Person]()
	if err = list.Append(Person{Name: "Alice", Age: 30}); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append(Person{Name: "Bob", Age: 25}); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	val, _ := list.Get(0)
	if val.Name != "Alice" || val.Age != 30 {
		t.Errorf("Get(0) = %+v, expected Alice/30", val)
	}

	// Out-of-range index returns an error (no wrapping).
	if _, err = list.Get(2); err == nil {
		t.Error("Get(2) on a 2-element list should return an out-of-range error")
	}
}

func TestCircularLinkedList_Set_OutOfRange(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if err := list.AppendAll(0, 1, 2); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Out-of-range indices return an error (no silent wrapping).
	if err := list.Set(3, 100); err == nil {
		t.Error("Set(3) on a 3-element list should return an out-of-range error")
	}
	if err := list.Set(-1, 200); err == nil {
		t.Error("Set(-1) should return an out-of-range error")
	}
}

// --- Tests for Head / PopHead (added 2026-07-14) ---

func TestCircularLinkedList_Head(t *testing.T) {
	list := NewCircularLinkedList[int]()

	if _, err := list.Head(); err == nil {
		t.Error("Head on an empty list should return an error")
	}

	if err := list.AppendAll(10, 20, 30); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	val, err := list.Head()
	if err != nil {
		t.Errorf("Head failed: %s", err)
	}
	if val != 10 {
		t.Errorf("Head = %d, expected 10", val)
	}

	// Head is a peek: it must not mutate the list.
	if list.Size() != 3 {
		t.Errorf("Head should not change size, got %d, expected 3", list.Size())
	}
}

func TestCircularLinkedList_PopHead(t *testing.T) {
	list := NewCircularLinkedList[int]()

	if _, err := list.PopHead(); err == nil {
		t.Error("PopHead on an empty list should return an error")
	}

	if err := list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	val, err := list.PopHead()
	if err != nil {
		t.Errorf("PopHead failed: %s", err)
	}
	if val != 1 {
		t.Errorf("PopHead = %d, expected 1", val)
	}
	if list.Size() != 2 {
		t.Errorf("Size after PopHead = %d, expected 2", list.Size())
	}
	if list.head.Data != 2 {
		t.Errorf("New head data = %d, expected 2", list.head.Data)
	}

	// The ring must remain closed: tail.Next should point to the new head.
	if list.tail.Next != list.head {
		t.Error("After PopHead, tail.Next should point back to the new head")
	}
}

func TestCircularLinkedList_PopHead_DrainOrder(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if err := list.AppendAll(1, 2, 3, 4, 5); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	for _, want := range []int{1, 2, 3, 4, 5} {
		got, err := list.PopHead()
		if err != nil {
			t.Fatalf("PopHead failed: %s", err)
		}
		if got != want {
			t.Errorf("PopHead = %d, expected %d", got, want)
		}
	}
	if list.Size() != 0 {
		t.Errorf("List should be empty and return its size to 0, but got = %d", list.Size())
	}
	if !list.IsEmpty() {
		t.Errorf("List should be empty after draining, size = %d", list.Size())
	}
	if _, err := list.PopHead(); err == nil {
		t.Error("PopHead on a drained list should return an error")
	}
}

func TestCircularLinkedList_PopHead_SingleElement(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PopHead on a single-element list should not panic: %v", r)
		}
	}()

	list := NewCircularLinkedList[int]()
	if err := list.Append(42); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	val, err := list.PopHead()
	if err != nil {
		t.Errorf("PopHead failed: %s", err)
	}
	if val != 42 {
		t.Errorf("PopHead = %d, expected 42", val)
	}
	if !list.IsEmpty() {
		t.Errorf("List should be empty after popping the only element, size = %d", list.Size())
	}

	// Once empty, Head and PopHead must both signal emptiness.
	if _, err := list.Head(); err == nil {
		t.Error("Head on the now-empty list should return an error")
	}
	if _, err := list.PopHead(); err == nil {
		t.Error("PopHead on the now-empty list should return an error")
	}
}
