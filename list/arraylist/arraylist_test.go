package arraylist

import "testing"

func TestArrayList_New(t *testing.T) {
	list := New[int](10)

	if list.Size() != 0 {
		t.Errorf("New list should have size 0, got %d", list.Size())
	}
	if list.listCapacity != 10 {
		t.Errorf("New list should have capacity 10, got %d", list.listCapacity)
	}
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
}

func TestArrayList_Append(t *testing.T) {
	testCases := []struct {
		name         string
		capacity     int
		appendCount  int
		expectedSize int
	}{
		{
			name:         "Append 10 elements with capacity 10",
			capacity:     10,
			appendCount:  10,
			expectedSize: 10,
		},
		{
			name:         "Append 11 elements with capacity 10 (triggers resize)",
			capacity:     10,
			appendCount:  11,
			expectedSize: 11,
		},
		{
			name:         "Append to zero capacity list",
			capacity:     0,
			appendCount:  5,
			expectedSize: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New[int](tc.capacity)

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

func TestArrayList_AppendAll(t *testing.T) {
	list := New[int](5)

	err := list.AppendAll(1, 2, 3, 4, 5)
	if err != nil {
		t.Errorf("AppendAll failed: %s", err)
	}

	if list.Size() != 5 {
		t.Errorf("Expected size 5, got %d", list.Size())
	}

	// Append more to trigger resize
	err = list.AppendAll(6, 7, 8)
	if err != nil {
		t.Errorf("AppendAll (resize) failed: %s", err)
	}

	if list.Size() != 8 {
		t.Errorf("Expected size 8, got %d", list.Size())
	}
}

func TestArrayList_Add(t *testing.T) {
	list := New[int](10)
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

	// Insert at the end
	err = list.Add(list.Size(), 5)
	if err != nil {
		t.Errorf("Add at end failed: %s", err)
	}
	val, _ = list.Get(list.Size() - 1)
	if val != 5 {
		t.Errorf("Get(last) = %d, expected 5", val)
	}
}

func TestArrayList_Add_OutOfBounds(t *testing.T) {
	list := New[int](10)
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
		t.Error("Add with index > size should return error")
	}
}

func TestArrayList_Set(t *testing.T) {
	list := New[int](10)
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

func TestArrayList_Get_OutOfBounds(t *testing.T) {
	list := New[int](10)
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

func TestArrayList_Delete(t *testing.T) {
	list := New[int](10)
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

	// Delete last element
	if err = list.Delete(list.Size() - 1); err != nil {
		t.Errorf("Delete failed: %s", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestArrayList_Delete_OutOfBounds(t *testing.T) {
	list := New[int](10)
	var err error
	if err = list.Append(1); err != nil {
		t.Errorf("Append failed: %s", err)
	}

	if err = list.Delete(-1); err == nil {
		t.Errorf("Delete with index <= 0 should return error")
	}

	if err = list.Delete(1); err == nil {
		t.Error("Delete with index >= size should return error")
	}
}

func TestArrayList_Clear(t *testing.T) {
	list := New[int](10)
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
}

func TestArrayList_CapacityGrowth(t *testing.T) {
	list := New[int](2)

	// Append more than initial capacity
	for i := 0; i < 100; i++ {
		err := list.Append(i)
		if err != nil {
			t.Errorf("Append failed at %d: %s", i, err)
		}
	}

	if list.Size() != 100 {
		t.Errorf("Expected size 100, got %d", list.Size())
	}

	// Verify all elements
	for i := 0; i < 100; i++ {
		val, err := list.Get(i)
		if err != nil {
			t.Errorf("Get(%d) failed: %s", i, err)
		}
		if val != i {
			t.Errorf("Get(%d) = %d, expected %d", i, val, i)
		}
	}
}

func TestArrayList_CapacityShrink(t *testing.T) {
	list := New[int](100)

	// Fill the list
	var err error
	for i := 0; i < 100; i++ {
		if err = list.Append(i); err != nil {
			t.Errorf("Append failed at %d: %s", i, err)
		}
	}

	initialCapacity := list.listCapacity

	// Delete most elements to trigger shrink
	for list.Size() > 10 {
		if err = list.Delete(0); err != nil {
			t.Errorf("Delete failed: %s", err)
		}
	}

	if list.listCapacity >= initialCapacity {
		t.Error("Capacity should have shrunk after deleting many elements")
	}

	// Verify the remaining elements are correct
	if list.Size() != 10 {
		t.Errorf("Expected size 10, got %d", list.Size())
	}
}

func TestArrayList_WithStrings(t *testing.T) {
	list := New[string](5)

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

func TestArrayList_WithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	var err error
	list := New[Person](5)
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
