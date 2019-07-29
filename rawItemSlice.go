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
	Set() *RawItemSlice
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

// Append adds a new raw Item pointer to the end of raw Item slice and returns the modified raw Item slice.
func (pointer *RawItemSlice) Append(rawItem *RawItem) *RawItemSlice {
	pointer.slice.Append(rawItem)
	return pointer
}

// Assign adds N raw Item pointers to the end raw Item slice and returns the modified raw Item slice.
func (pointer *RawItemSlice) Assign(rawItem ...*RawItem) *RawItemSlice {
	for _, rawItem := range rawItem {
		pointer.Append(rawItem)
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the raw Item slice.
func (pointer *RawItemSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two raw Item slices into a single raw Item slice.
func (pointer *RawItemSlice) Concatenate(rawItemSlice *RawItemSlice) *RawItemSlice {
	pointer.slice.Concatenate(rawItemSlice.slice)
	return pointer
}

// Each executes a provided function once for each element in the raw Item slice.
func (pointer *RawItemSlice) Each(f func(i int, rawItem *RawItem)) *RawItemSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*RawItem))
	})
	return pointer
}

// Empty returns a boolean indicating whether the raw Item slice contains zero values.
func (pointer *RawItemSlice) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the raw Item pointer held at the argument index. Returns nil if index exceeds raw Item slice length.
func (pointer *RawItemSlice) Fetch(i int) *RawItem {
	rawItem, _ := pointer.Get(i)
	return rawItem
}

// Get returns the raw Item pointer held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *RawItemSlice) Get(i int) (*RawItem, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*RawItem), ok
	}
	return nil, ok
}

// Len method returns the number of elements in the raw Item slice.
func (pointer *RawItemSlice) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each raw Item pointer in the raw Item slice
// and sets the returned value to the current index.
func (pointer *RawItemSlice) Map(f func(i int, rawItem *RawItem) *RawItem) *RawItemSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*RawItem))
	})
	return pointer
}

// Poll method removes the first raw Item pointer from the raw Item slice and returns that removed pointer.
func (pointer *RawItemSlice) Poll() *RawItem {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawItem)
	}
	return nil
}

// Pop method removes the last raw Item from the raw Item slice and returns that pointer.
func (pointer *RawItemSlice) Pop() *RawItem {
	value := pointer.slice.Pop()
	if value != nil {
		return value.(*RawItem)
	}
	return nil
}

// Preassign method adds zero or more raw Item pointers to the beginning of the raw Item slice and returns the modified raw Item slice.
func (pointer *RawItemSlice) Preassign(rawItem ...*RawItem) *RawItemSlice {
	for _, rawItem := range rawItem {
		pointer.Prepend(rawItem)
	}
	return pointer
}

// Precatenate merges two raw Item slices, prepending the argument raw Item slice to the beginning of the receiver raw Item slice.
func (pointer *RawItemSlice) Precatenate(rawItemSlice *RawItemSlice) *RawItemSlice {
	pointer.slice.Precatenate(rawItemSlice.slice)
	return pointer
}

// Prepend method adds one raw Item to the beginning of the raw Item sclie and returns the modified raw Item slice.
func (pointer *RawItemSlice) Prepend(rawItem *RawItem) *RawItemSlice {
	pointer.slice.Prepend(rawItem)
	return pointer
}

// Push method adds a new raw Item to the end of the raw Item slice and returns the length of the modified raw Item slice.
func (pointer *RawItemSlice) Push(rawItem *RawItem) int {
	return pointer.slice.Push(rawItem)
}

// Replace method replaces the raw Item at the argument index if it is in bounds with the provided argument raw Item.
func (pointer *RawItemSlice) Replace(i int, rawItem *RawItem) bool {
	return pointer.slice.Replace(i, rawItem)
}

// Set method returns a unique raw Item slice, removing duplicate raw Item pointers that have the same title.
func (pointer *RawItemSlice) Set() *RawItemSlice {
	rawItemSlice := newRawItemSlice()
	m := map[string]bool{}
	pointer.Each(func(_ int, rawItem *RawItem) {
		if _, ok := m[rawItem.Title]; !ok {
			m[rawItem.Title] = true
			rawItemSlice.Append(rawItem)
		}
	})
	return rawItemSlice
}

// Slice method returns a shallow copy of a portion of the raw Item slice into a new raw Item slice.
// Raw Item slice is selected from begin to end (end not included).
// The original raw Item slice will not be modified but all values are shared between the two raw Item slices.
func (pointer *RawItemSlice) Slice(start, end int) *RawItemSlice {
	return &RawItemSlice{slice: pointer.slice.Slice(start, end)}
}

// Splice method changes the contents of the raw Item slice by removing existing elements fron i to N.
// Returns a new raw Item slice containing the cut values.
func (pointer *RawItemSlice) Splice(start, end int) *RawItemSlice {
	return &RawItemSlice{slice: pointer.slice.Splice(start, end)}
}

// String returns the string value of the raw Item slice.
func (pointer *RawItemSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
