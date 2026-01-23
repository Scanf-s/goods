package arraylist

import "list"

type Arraylist interface {
	EnsureCapacity(list *list.List)
}

func Append()