package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ rawItemSlice = (*RawItemSlice)(nil)
)

func getRawItemSlice(content *[]byte) *RawItemSlice {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	rawItemSlice := newRawItemSlice()
	for _, rawMessage := range rawPayload.Items {
		rawItemSlice.Append(newRawItem(rawMessage))
	}
	return rawItemSlice
}

func newRawItemSlice() *RawItemSlice {
	return &RawItemSlice{slice: &slice.Slice{}}
}

// rawItemSlice defines the interface for an raw Item slice pointer.
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

// An RawItemSlice is a slice-like struct whose methods are used to perform insertion, mutation and iteration operations on an
// unordered collection of raw Item pointers. Each raw Item slice can contain 0 to N number of raw Items, with each
// raw Item pointer being held in a private slice field. All exposed methods for the raw Item slice perform a corresponding
// operation for this internal field. This property is protected to prevent incorrect data assignment as the slice permits
// any data interface to be assigned to the raw Item slice. Raw Item slices contain the as-is provided
// raw Item collected from the Nintendo XHR HTTP response.
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
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawItem)
	}
	return nil
}

func (pointer *RawItemSlice) Pop() *RawItem {
	value := pointer.slice.Pop()
	if value != nil {
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
