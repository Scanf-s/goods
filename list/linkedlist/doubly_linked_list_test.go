package linkedlist

import "testing"

func TestDoublyLinkedList_New(t *testing.T) {
	list := NewDoublyLinkedList[int]()

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

func TestDoublyLinkedList_Append(t *testing.T) {
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
			list := NewDoublyLinkedList[int]()

			for i := 0; i < tc.appendCount; i++ {
				err := list.Append(i)
				if err != nil {
					t.Errorf("Append failed: %s", err)
				}
			}

			if list.Size() != tc.expectedSize {
				t.Errorf("Expected size %d, got %d", tc.expectedSize, list.Size())
			}

			// Verify elements are stored correctly
			for i := 0; i < tc.appendCount; i++ {
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

func TestDoublyLinkedList_Append_HeadTail(t *testing.T) {
	list := NewDoublyLinkedList[int]()

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

	// Verify Prev pointer
	if list.tail.Prev != list.head {
		t.Error("Tail's Prev should point to head")
	}
	if list.head.Prev != nil {
		t.Error("Head's Prev should be nil")
	}
}

func TestDoublyLinkedList_Append_PrevPointers(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	if err := list.AppendAll(1, 2, 3, 4, 5); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Traverse forward and collect values
	forward := make([]int, 0, 5)
	current := list.head
	for current != nil {
		forward = append(forward, current.Data)
		current = current.Next
	}

	// Traverse backward and collect values
	backward := make([]int, 0, 5)
	current = list.tail
	for current != nil {
		backward = append(backward, current.Data)
		current = current.Prev
	}

	// Verify forward traversal
	expectedForward := []int{1, 2, 3, 4, 5}
	for i, exp := range expectedForward {
		if forward[i] != exp {
			t.Errorf("Forward[%d] = %d, expected %d", i, forward[i], exp)
		}
	}

	// Verify backward traversal
	expectedBackward := []int{5, 4, 3, 2, 1}
	for i, exp := range expectedBackward {
		if backward[i] != exp {
			t.Errorf("Backward[%d] = %d, expected %d", i, backward[i], exp)
		}
	}
}

func TestDoublyLinkedList_AppendAll(t *testing.T) {
	list := NewDoublyLinkedList[int]()

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
}

func TestDoublyLinkedList_Add(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

	// Insert at the beginning
	err = list.Add(0, 0)
	if err != nil {
		t.Errorf("Add at beginning failed: %s", err)
	}
	val, _ := list.Get(0)
	if val != 0 {
		t.Errorf("Get(0) = %d, expected 0", val)
	}
	if list.head.Data != 0 {
		t.Errorf("Head should be updated to 0, got %d", list.head.Data)
	}
	if list.head.Prev != nil {
		t.Error("New head's Prev should be nil")
	}
}

func TestDoublyLinkedList_Add_VerifyPrevPointers(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Insert in middle
	if err = list.Add(1, 2); err != nil {
		t.Errorf("Add failed: %s", err)
	}
	// list: [1, 2, 3]

	// Verify forward and backward links
	node1 := list.head
	node2 := node1.Next
	node3 := node2.Next

	if node2.Prev != node1 {
		t.Error("Node2's Prev should point to Node1")
	}
	if node3.Prev != node2 {
		t.Error("Node3's Prev should point to Node2")
	}
	if node1.Next != node2 {
		t.Error("Node1's Next should point to Node2")
	}
	if node2.Next != node3 {
		t.Error("Node2's Next should point to Node3")
	}
}

func TestDoublyLinkedList_Add_OutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	err := list.Append(1)
	if err != nil {
		t.Errorf("Append failed: %s", err)
	}

	err = list.Add(-1, 0)
	if err == nil {
		t.Error("Add with negative index should return error")
	}

	err = list.Add(10, 0)
	if err == nil {
		t.Error("Add with index >= size should return error")
	}
}

func TestDoublyLinkedList_Set(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append(2); err != nil {
		t.Errorf("Append failed: %s", err)
	}
	if err = list.Append(3); err != nil {
		t.Errorf("Append failed: %s", err)
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

func TestDoublyLinkedList_Set_OutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	err = list.Set(-1, 0)
	if err == nil {
		t.Error("Set with negative index should return error")
	}

	err = list.Set(1, 0)
	if err == nil {
		t.Error("Set with index >= size should return error")
	}
}

func TestDoublyLinkedList_Get_OutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	_, err = list.Get(-1)
	if err == nil {
		t.Error("Get with negative index should return error")
	}

	_, err = list.Get(1)
	if err == nil {
		t.Error("Get with index >= size should return error")
	}

	_, err = list.Get(100)
	if err == nil {
		t.Error("Get with index >= size should return error")
	}
}

func TestDoublyLinkedList_Delete(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

	// Delete first element
	if err = list.Delete(0); err != nil {
		t.Errorf("Delete failed: %s", err)
	}
	val, _ := list.Get(0)
	if val != 2 {
		t.Errorf("After deleting first, Get(0) = %d, expected 2", val)
	}
	if list.head.Data != 2 {
		t.Errorf("Head should be updated to 2, got %d", list.head.Data)
	}
	if list.head.Prev != nil {
		t.Error("New head's Prev should be nil")
	}

	// Delete last element
	if err = list.Delete(list.Size() - 1); err != nil {
		t.Errorf("Delete failed: %s", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	if list.tail.Data != 4 {
		t.Errorf("Tail should be updated to 4, got %d", list.tail.Data)
	}
}

func TestDoublyLinkedList_Delete_VerifyPrevPointers(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.AppendAll(1, 2, 3); err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	// Delete middle
	if err = list.Delete(1); err != nil {
		t.Errorf("Delete failed: %s", err)
	}
	// list: [1, 3]

	node1 := list.head
	node3 := list.tail

	if node1.Next != node3 {
		t.Error("After delete, Node1's Next should point to Node3")
	}
	if node3.Prev != node1 {
		t.Error("After delete, Node3's Prev should point to Node1")
	}
}

func TestDoublyLinkedList_Delete_SingleElement(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDoublyLinkedList_Delete_OutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	if err = list.Delete(-1); err == nil {
		t.Error("Delete with negative index should return error")
	}

	if err = list.Delete(1); err == nil {
		t.Error("Delete with index >= size should return error")
	}
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDoublyLinkedList_WithStrings(t *testing.T) {
	list := NewDoublyLinkedList[string]()

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
}

func TestDoublyLinkedList_WithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	var err error
	list := NewDoublyLinkedList[Person]()
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
}
