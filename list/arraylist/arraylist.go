package arraylist

import "goods/list"

type Arraylist interface {
	EnsureCapacity(list *list.List)
}

func Append()