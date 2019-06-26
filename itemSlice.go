package main

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ amiiboItemSlice = (*AmiiboItemSlice)(nil)
)

type amiiboItemSlice interface {
	Append(amiiboItem *AmiiboItem) *AmiiboItemSlice
	Assign(amiiboItem ...*AmiiboItem) *AmiiboItemSlice
	Bounds(i int) bool
	Concatenate(amiiboItemSlice *AmiiboItemSlice) *AmiiboItemSlice
	Each(f func(i int, amiiboItem *AmiiboItem)) *AmiiboItemSlice
	Empty() bool
	Fetch(i int) *AmiiboItem
	Get(i int) (*AmiiboItem, bool)
	Len() int
	Map(func(i int, amiiboItem *AmiiboItem) *AmiiboItem) *AmiiboItemSlice
	Poll() *AmiiboItem
	Pop() *AmiiboItem
	Preassign(amiiboItem ...*AmiiboItem) *AmiiboItemSlice
	Precatenate(amiiboItemSlice *AmiiboItemSlice) *AmiiboItemSlice
	Prepend(amiiboItem *AmiiboItem) *AmiiboItemSlice
	Push(amiiboItem *AmiiboItem) int
	Replace(i int, amiiboItem *AmiiboItem) bool
	Slice(start, end int) *AmiiboItemSlice
	Splice(start, end int) *AmiiboItemSlice
	String() string
}

type AmiiboItemSlice struct {
	slice *slice.Slice
}

func (pointer *AmiiboItemSlice) Append(amiiboItem *AmiiboItem) *AmiiboItemSlice {
	pointer.slice.Append(amiiboItem)
	return pointer
}

func (pointer *AmiiboItemSlice) Assign(amiiboItem ...*AmiiboItem) *AmiiboItemSlice {
	for _, amiiboItem := range amiiboItem {
		pointer.Append(amiiboItem)
	}
	return pointer
}

func (pointer *AmiiboItemSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *AmiiboItemSlice) Concatenate(amiiboItemSlice *AmiiboItemSlice) *AmiiboItemSlice {
	pointer.slice.Concatenate(amiiboItemSlice.slice)
	return pointer
}

func (pointer *AmiiboItemSlice) Each(f func(i int, amiiboItem *AmiiboItem)) *AmiiboItemSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*AmiiboItem))
	})
	return pointer
}

func (pointer *AmiiboItemSlice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *AmiiboItemSlice) Fetch(i int) *AmiiboItem {
	amiibo, _ := pointer.Get(i)
	return amiibo
}

func (pointer *AmiiboItemSlice) Get(i int) (*AmiiboItem, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*AmiiboItem), ok
	}
	return nil, ok
}

func (pointer *AmiiboItemSlice) Len() int {
	return pointer.slice.Len()
}

func (pointer *AmiiboItemSlice) Map(f func(i int, amiiboItem *AmiiboItem) *AmiiboItem) *AmiiboItemSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*AmiiboItem))
	})
	return pointer
}

func (pointer *AmiiboItemSlice) Poll() *AmiiboItem {
	if value := pointer.slice.Poll(); value != nil {
		return value.(*AmiiboItem)
	}
	return nil
}

func (pointer *AmiiboItemSlice) Pop() *AmiiboItem {
	if value := pointer.slice.Pop(); value != nil {
		return value.(*AmiiboItem)
	}
	return nil
}

func (pointer *AmiiboItemSlice) Preassign(amiiboItem ...*AmiiboItem) *AmiiboItemSlice {
	for _, amiiboItem := range amiiboItem {
		pointer.Prepend(amiiboItem)
	}
	return pointer
}

func (pointer *AmiiboItemSlice) Precatenate(amiiboItemSlice *AmiiboItemSlice) *AmiiboItemSlice {
	pointer.slice.Precatenate(amiiboItemSlice.slice)
	return pointer
}

func (pointer *AmiiboItemSlice) Prepend(amiiboItem *AmiiboItem) *AmiiboItemSlice {
	pointer.slice.Prepend(amiiboItem)
	return pointer
}

func (pointer *AmiiboItemSlice) Push(amiiboItem *AmiiboItem) int {
	return pointer.slice.Push(amiiboItem)
}

func (pointer *AmiiboItemSlice) Replace(i int, amiiboItem *AmiiboItem) bool {
	return pointer.slice.Replace(i, amiiboItem)
}

func (pointer *AmiiboItemSlice) Slice(start, end int) *AmiiboItemSlice {
	return &AmiiboItemSlice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *AmiiboItemSlice) Splice(start, end int) *AmiiboItemSlice {
	return &AmiiboItemSlice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *AmiiboItemSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
