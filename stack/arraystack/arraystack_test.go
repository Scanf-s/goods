package arraystack

import "testing"

func TestArrayStack_New(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
}

func TestArrayStack_Push_Integers(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	arrayStack.Push(1)
	arrayStack.Push(2)
	arrayStack.Push(3)

	if arrayStack.Size() != 3 {
		t.Errorf("Stack should have 3 elements but got %d", arrayStack.Size())
	}
	if top, err := arrayStack.Top(); err != nil && top != 3 {
		t.Errorf("Stack's top should be 3 without any error")
	}
}

func TestArrayStack_Push_Objects(t *testing.T) {
	type Student struct {
		name string
		gpa float32
	}
	arrayStack := NewArrayStack[Student]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	arrayStack.Push(Student{name:"test1", gpa:3.2})
	arrayStack.Push(Student{name:"test2", gpa:3.5})
	arrayStack.Push(Student{name:"test3", gpa:4.0})

	if arrayStack.Size() != 3 {
		t.Errorf("Stack should have 3 elements but got %d", arrayStack.Size())
	}
	if top, err := arrayStack.Top(); err != nil && top.name != "test3" {
		t.Errorf("Stack's top should be student with name `test3` without any error")
	}
}

func TestArrayStack_Pop_NormalCase(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}

	arrayStack.Push(1)
	arrayStack.Push(2)
	arrayStack.Push(3)

	for i := 3; i > 0; i-- {
		num, err := arrayStack.Pop()
		if err != nil {
			t.Errorf("Stack pop operation should succeed without an error %v", err)
		}
		if num != i {
			t.Errorf("Stack should returns the top element, But got %d", num)
		}
	}

	if arrayStack.Size() != 0 {
		t.Errorf("Stack should be empty")
	}
}

func TestArrayStack_Pop_EmptyStack(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}

	if _, err := arrayStack.Pop(); err == nil {
		t.Errorf("Stack should return an error if user is trying to pop the empty stack")
	}
}

func TestArrayStack_Top_ShouldReturnValue(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	arrayStack.Push(1)

	if num, err := arrayStack.Top(); err != nil || num != 1 {
		t.Errorf("Stack.Top() should return value 1 without an error")
	}
}

func TestArrayStack_Top_ShouldReturnError(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}

	if _, err := arrayStack.Top(); err == nil {
		t.Errorf("Stack.Top() should return an error if user is trying to fetch top element of the empty stack")
	}	
}

func TestArrayStack_IsEmpty_True(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	
	if !arrayStack.IsEmpty() {
		t.Error("Stack should be empty")
	}
}

func TestArrayStack_IsEmpty_False(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	arrayStack.Push(1)
	
	if arrayStack.IsEmpty() {
		t.Error("Stack should be not empty")
	}
}

func TestArrayStack_Size(t *testing.T) {
	arrayStack := NewArrayStack[int]()
	if arrayStack == nil {
		t.Errorf("New stack array has to be initialized")
	}
	arrayStack.Push(1)
	
	if arrayStack.Size() != 1 {
		t.Error("Stack size should be 1")
	}
}