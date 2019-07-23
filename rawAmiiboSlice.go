package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ rawAmiiboSlice = (*RawAmiiboSlice)(nil)
)

// getRawAmiiboSlice returns a new raw Amiibo slice pointer using the argument bytes as the initial entries.
func getRawAmiiboSlice(content *[]byte) *RawAmiiboSlice {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	rawAmiiboSlice := newRawAmiiboSlice()
	for _, rawMessage := range rawPayload.AmiiboList {
		rawAmiiboSlice.Append(newRawAmiibo(rawMessage))
	}
	return rawAmiiboSlice
}

// newRawAmiiboSlice instantiates a new raw Amiibo slice pointer.
func newRawAmiiboSlice() *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: &slice.Slice{}}
}

// rawAmiiboSlice defines the interface for a raw Amiibo slice pointer.
type rawAmiiboSlice interface {
	Append(rawAmiibo *RawAmiibo) *RawAmiiboSlice
	Assign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice
	Bounds(i int) bool
	Concatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice
	Each(f func(i int, rawAmiibo *RawAmiibo)) *RawAmiiboSlice
	Empty() bool
	Fetch(i int) *RawAmiibo
	Get(i int) (*RawAmiibo, bool)
	Len() int
	Map(func(i int, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboSlice
	Poll() *RawAmiibo
	Pop() *RawAmiibo
	Preassign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice
	Precatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice
	Prepend(rawAmiibo *RawAmiibo) *RawAmiiboSlice
	Push(rawAmiibo *RawAmiibo) int
	Replace(i int, rawAmiibo *RawAmiibo) bool
	Slice(start, end int) *RawAmiiboSlice
	Splice(start, end int) *RawAmiiboSlice
	String() string
}

// An RawAmiiboSlice is a slice-like struct whose methods are used to perform insertion, mutation and iteration operations on an
// unordered collection of raw Amiibo pointers. Each raw Amiibo slice can contain 0 to N number of raw Amiibo, with each
// raw Amiibo pointer being held in a private slice field. All exposed methods for the raw Amiibo slice perform a corresponding
// operation for this internal field. This property is protected to prevent incorrect data assignment as the slice permits
// any data interface to be assigned to the raw Amiibo slice. Raw Amiibo slices contain the as-is provided
// raw Amiibo collected from the Nintendo XHR HTTP response.
type RawAmiiboSlice struct {
	slice *slice.Slice
}

// Append adds a new raw Amiibo to the end of raw Amiibo slice and returns the modified raw Amiibo slice.
func (pointer *RawAmiiboSlice) Append(rawAmiibo *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Append(rawAmiibo)
	return pointer
}

// Assign adds N raw Amiibo to the end raw Amiibo slice and returns the modified raw Amiibo slice.
func (pointer *RawAmiiboSlice) Assign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice {
	for _, rawAmiibo := range rawAmiibo {
		pointer.Append(rawAmiibo)
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the raw Amiibo slice.
func (pointer *RawAmiiboSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two raw Amiibo slices into a single raw Amiibo slice.
func (pointer *RawAmiiboSlice) Concatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice {
	pointer.slice.Concatenate(rawAmiiboSlice.slice)
	return pointer
}

// Each executes a provided function once for each element in the raw Amiibo slice.
func (pointer *RawAmiiboSlice) Each(f func(i int, rawAmiibo *RawAmiibo)) *RawAmiiboSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*RawAmiibo))
	})
	return pointer
}

// Empty returns a boolean indicating whether the raw Amiibo slice contains zero values.
func (pointer *RawAmiiboSlice) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the raw Amiibo pointer held at the argument index. Returns nil if index exceeds raw Amiibo slice length.
func (pointer *RawAmiiboSlice) Fetch(i int) *RawAmiibo {
	rawAmiibo, _ := pointer.Get(i)
	return rawAmiibo
}

// Get returns the raw Amiibo pointer held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *RawAmiiboSlice) Get(i int) (*RawAmiibo, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*RawAmiibo), ok
	}
	return nil, ok
}

// Len method returns the number of elements in the raw Amiibo slice.
func (pointer *RawAmiiboSlice) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each raw Amiibo pointer in the raw Amiibo slice
// and sets the returned value to the current index.
func (pointer *RawAmiiboSlice) Map(f func(i int, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*RawAmiibo))
	})
	return pointer
}

// Poll method removes the first raw Amiibo pointer from the raw Amiibo slice and returns that removed pointer.
func (pointer *RawAmiiboSlice) Poll() *RawAmiibo {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawAmiibo)
	}
	return nil
}

// Pop method removes the last raw Amiibo from the raw Amiibo slice and returns that pointer.
func (pointer *RawAmiiboSlice) Pop() *RawAmiibo {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawAmiibo)
	}
	return nil
}

// Preassign method adds zero or more raw Amiibo pointers to the beginning of the raw Amiibo slice and returns the modified raw Amiibo slice.
func (pointer *RawAmiiboSlice) Preassign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice {
	for _, rawAmiibo := range rawAmiibo {
		pointer.Prepend(rawAmiibo)
	}
	return pointer
}

// Precatenate merges two raw Amiibo slices, prepending the argument raw Amiibo slice to the beginning of the receiver raw Amiibo slice.
func (pointer *RawAmiiboSlice) Precatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice {
	pointer.slice.Precatenate(rawAmiiboSlice.slice)
	return pointer
}

// Prepend method adds one raw Amiibo to the beginning of the raw Amiibo sclie and returns the modified raw Amiibo slice.
func (pointer *RawAmiiboSlice) Prepend(rawAmiibo *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Prepend(rawAmiibo)
	return pointer
}

// Push method adds a new raw Amiibo to the end of the raw Amiibo slice and returns the length of the modified raw Amiibo slice.
func (pointer *RawAmiiboSlice) Push(rawAmiibo *RawAmiibo) int {
	return pointer.slice.Push(rawAmiibo)
}

// Replace method replaces the raw Amiibo at the argument index if it is in bounds with the provided argument raw Amiibo.
func (pointer *RawAmiiboSlice) Replace(i int, rawAmiibo *RawAmiibo) bool {
	return pointer.slice.Replace(i, rawAmiibo)
}

// Set method returns a unique raw Amiibo slice, removing duplicate raw Amiibo that have the same name.
func (pointer *RawAmiiboSlice) Set() *RawAmiiboSlice {
	rawAmiiboSlice := newRawAmiiboSlice()
	m := map[string]bool{}
	pointer.Each(func(_ int, rawAmiibo *RawAmiibo) {
		if _, ok := m[rawAmiibo.AmiiboName]; !ok {
			m[rawAmiibo.AmiiboName] = true
			rawAmiiboSlice.Append(rawAmiibo)
		}
	})
	return rawAmiiboSlice
}

// Slice method returns a shallow copy of a portion of the raw Amiibo slice into a new raw Amiibo slice.
// Raw Amiibo slice is selected from begin to end (end not included).
// The original raw Amiibo slice will not be modified but all values are shared between the two raw Amiibo slices.
func (pointer *RawAmiiboSlice) Slice(start, end int) *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: pointer.slice.Slice(start, end)}
}

// Splice method changes the contents of the raw Amiibo slice by removing existing elements fron i to N.
// Returns a new raw Amiibo slice containing the cut values.
func (pointer *RawAmiiboSlice) Splice(start, end int) *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: pointer.slice.Splice(start, end)}
}

// String returns the string value of the raw Amiibo slice.
func (pointer *RawAmiiboSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
