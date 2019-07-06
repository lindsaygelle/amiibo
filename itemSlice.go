package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ itemSlice = (*ItemSlice)(nil)
)

func getItemSlice(content *[]byte) *ItemSlice {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	itemSlice := newItemSlice()
	for _, rawMessage := range rawPayload.Items {
		r := newRawItem(rawMessage)
		a := newItem(r)
		itemSlice.Append(a)
	}
	return itemSlice
}

func newItemSlice() *ItemSlice {
	return &ItemSlice{slice: &slice.Slice{}}
}

type itemSlice interface {
	Append(item *Item) *ItemSlice
	Assign(item ...*Item) *ItemSlice
	Bounds(i int) bool
	Concatenate(itemSlice *ItemSlice) *ItemSlice
	Each(f func(i int, item *Item)) *ItemSlice
	Empty() bool
	Fetch(i int) *Item
	Get(i int) (*Item, bool)
	Len() int
	Map(func(i int, item *Item) *Item) *ItemSlice
	Poll() *Item
	Pop() *Item
	Preassign(item ...*Item) *ItemSlice
	Precatenate(itemSlice *ItemSlice) *ItemSlice
	Prepend(item *Item) *ItemSlice
	Push(item *Item) int
	Replace(i int, item *Item) bool
	Slice(start, end int) *ItemSlice
	Splice(start, end int) *ItemSlice
	String() string
}

type ItemSlice struct {
	slice *slice.Slice
}

func (pointer *ItemSlice) Append(item *Item) *ItemSlice {
	pointer.slice.Append(item)
	return pointer
}

func (pointer *ItemSlice) Assign(item ...*Item) *ItemSlice {
	for _, item := range item {
		pointer.Append(item)
	}
	return pointer
}

func (pointer *ItemSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *ItemSlice) Concatenate(itemSlice *ItemSlice) *ItemSlice {
	pointer.slice.Concatenate(itemSlice.slice)
	return pointer
}

func (pointer *ItemSlice) Each(f func(i int, item *Item)) *ItemSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Item))
	})
	return pointer
}

func (pointer *ItemSlice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *ItemSlice) Fetch(i int) *Item {
	item, _ := pointer.Get(i)
	return item
}

func (pointer *ItemSlice) Get(i int) (*Item, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Item), ok
	}
	return nil, ok
}

func (pointer *ItemSlice) Len() int {
	return pointer.slice.Len()
}

func (pointer *ItemSlice) Map(f func(i int, item *Item) *Item) *ItemSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Item))
	})
	return pointer
}

func (pointer *ItemSlice) Poll() *Item {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Item)
	}
	return nil
}

func (pointer *ItemSlice) Pop() *Item {
	value := pointer.slice.Pop()
	if value != nil {
		return value.(*Item)
	}
	return nil
}

func (pointer *ItemSlice) Preassign(item ...*Item) *ItemSlice {
	for _, item := range item {
		pointer.Prepend(item)
	}
	return pointer
}

func (pointer *ItemSlice) Precatenate(itemSlice *ItemSlice) *ItemSlice {
	pointer.slice.Precatenate(itemSlice.slice)
	return pointer
}

func (pointer *ItemSlice) Prepend(item *Item) *ItemSlice {
	pointer.slice.Prepend(item)
	return pointer
}

func (pointer *ItemSlice) Push(item *Item) int {
	return pointer.slice.Push(item)
}

func (pointer *ItemSlice) Replace(i int, item *Item) bool {
	return pointer.slice.Replace(i, item)
}

func (pointer *ItemSlice) Slice(start, end int) *ItemSlice {
	return &ItemSlice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *ItemSlice) Splice(start, end int) *ItemSlice {
	return &ItemSlice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *ItemSlice) String() string {
	return fmt.Sprintf("%v", *pointer)
}
