package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

func newRawItemSlice() *RawItemSlice {
	return new(RawItemSlice)
}

var (
	_ rawItemSlice = (*RawItemSlice)(nil)
)

type rawItemSlice interface {
	Append(rawItem *RawItem) *RawItemSlice
	Assign(rawItem ...*RawItem) *RawItemSlice
	Bounds(i int) bool
	Concatenate(rawItemSlice *RawItemSlice) *RawItemSlice
	Each(f func(i int, rawItem *RawItem)) *RawItemSlice
	Empty() bool
	Fetch(i int) *RawItem
	Get(i int) (*RawItem, bool)
	Len() int
	Map(func(i int, rawItem *RawItem) *RawItem) *RawItemSlice
	Poll() *RawItem
	Pop() *RawItem
	Preassign(rawItem ...*RawItem) *RawItemSlice
	Precatenate(rawItemSlice *RawItemSlice) *RawItemSlice
	Prepend(rawItem *RawItem) *RawItemSlice
	Push(rawItem *RawItem) int
	Replace(i int, rawItem *RawItem) bool
	Slice(start, end int) *RawItemSlice
	Splice(start, end int) *RawItemSlice
	String() string
}

type RawItemSlice struct {
	slice *slice.Slice
}

func (pointer *RawItemSlice) Append(rawItem *RawItem) *RawItemSlice {
	pointer.slice.Append(rawItem)
	return pointer
}

func (pointer *RawItemSlice) Assign(rawItem ...*RawItem) *RawItemSlice {
	for _, rawItem := range rawItem {
		pointer.Append(rawItem)
	}
	return pointer
}

func (pointer *RawItemSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *RawItemSlice) Concatenate(rawItemSlice *RawItemSlice) *RawItemSlice {
	pointer.slice.Concatenate(rawItemSlice.slice)
	return pointer
}

func (pointer *RawItemSlice) Each(f func(i int, rawItem *RawItem)) *RawItemSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*RawItem))
	})
	return pointer
}

func (pointer *RawItemSlice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *RawItemSlice) Fetch(i int) *RawItem {
	rawItem, _ := pointer.Get(i)
	return rawItem
}

func (pointer *RawItemSlice) Get(i int) (*RawItem, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*RawItem), ok
	}
	return nil, ok
}

func (pointer *RawItemSlice) Len() int {
	return pointer.slice.Len()
}

func (pointer *RawItemSlice) Map(f func(i int, rawItem *RawItem) *RawItem) *RawItemSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*RawItem))
	})
	return pointer
}

func (pointer *RawItemSlice) Poll() *RawItem {
	if value := pointer.slice.Poll(); value != nil {
		return value.(*RawItem)
	}
	return nil
}

func (pointer *RawItemSlice) Pop() *RawItem {
	if value := pointer.slice.Pop(); value != nil {
		return value.(*RawItem)
	}
	return nil
}

func (pointer *RawItemSlice) Preassign(rawItem ...*RawItem) *RawItemSlice {
	for _, rawItem := range rawItem {
		pointer.Prepend(rawItem)
	}
	return pointer
}

func (pointer *RawItemSlice) Precatenate(rawItemSlice *RawItemSlice) *RawItemSlice {
	pointer.slice.Precatenate(rawItemSlice.slice)
	return pointer
}

func (pointer *RawItemSlice) Prepend(rawItem *RawItem) *RawItemSlice {
	pointer.slice.Prepend(rawItem)
	return pointer
}

func (pointer *RawItemSlice) Push(rawItem *RawItem) int {
	return pointer.slice.Push(rawItem)
}

func (pointer *RawItemSlice) Replace(i int, rawItem *RawItem) bool {
	return pointer.slice.Replace(i, rawItem)
}

func (pointer *RawItemSlice) Slice(start, end int) *RawItemSlice {
	return &RawItemSlice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *RawItemSlice) Splice(start, end int) *RawItemSlice {
	return &RawItemSlice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *RawItemSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
