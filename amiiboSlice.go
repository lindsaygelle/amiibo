package amiibo

import (
	"encoding/json"
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ amiiboSlice = (*AmiiboSlice)(nil)
)

// NewAmiiboSlice returns a new Amiibo slice pointer. A Amiibo slice pointer can be built
// from a cached XHR payload or directly from the Nintendo Amiibo source. To create from source
// parse in the optional byte code pointer, otherwise leave empty and it will be collected from
// the Nintendo XHR HTTP response.
func NewAmiiboSlice(b ...byte) *AmiiboSlice {
	if len(b) != 0 {
		return getAmiiboSlice(&b)
	}
	x, err := net()
	if err != nil {
		return newAmiiboSlice()
	}
	r, err := unmarshallRawPayload(x)
	if err != nil {
		return newAmiiboSlice()
	}
	return unmarshallRawToAmiiboSlice(r.AmiiboList)
}

// getAmiiboSlice returns a new Amiibo slice pointer using the argument bytes as the initial entries.
func getAmiiboSlice(content *[]byte) *AmiiboSlice {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	amiiboSlice := newAmiiboSlice()
	for _, rawMessage := range rawPayload.AmiiboList {
		r := newRawAmiibo(rawMessage)
		a := newAmiibo(r)
		amiiboSlice.Append(a)
	}
	return amiiboSlice
}

// newAmiiboSlice instantiates a new Amiibo slice pointer.
func newAmiiboSlice() *AmiiboSlice {
	return &AmiiboSlice{slice: &slice.Slice{}}
}

// unmarshallRawToAmiiboSlice returns a Amiibo slice from the raw bytes contained within the Nintendo XHR HTTP response.
func unmarshallRawToAmiiboSlice(r []*json.RawMessage) *AmiiboSlice {
	amiiboSlice := newAmiiboSlice()
	for _, rawMessage := range r {
		amiiboSlice.Append(newAmiibo(newRawAmiibo(rawMessage)))
	}
	return amiiboSlice
}

// amiiboSlice defines the interface for an Amiibo slice pointer.
type amiiboSlice interface {
	Append(amiibo *Amiibo) *AmiiboSlice
	Assign(amiibo ...*Amiibo) *AmiiboSlice
	Bounds(i int) bool
	Concatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice
	Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice
	EachReverse(f func(i int, amiibo *Amiibo)) *AmiiboSlice
	Empty() bool
	Fetch(i int) *Amiibo
	Get(i int) (*Amiibo, bool)
	Len() int
	Map(func(i int, amiibo *Amiibo) *Amiibo) *AmiiboSlice
	Poll() *Amiibo
	Pop() *Amiibo
	Preassign(amiibo ...*Amiibo) *AmiiboSlice
	Precatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice
	Prepend(amiibo *Amiibo) *AmiiboSlice
	Push(amiibo *Amiibo) int
	Replace(i int, amiibo *Amiibo) bool
	Set() *AmiiboSlice
	Slice(start, end int) *AmiiboSlice
	Splice(start, end int) *AmiiboSlice
	String() string
	Swap(i, j int) *AmiiboSlice
}

// An AmiiboSlice is a slice-like struct whose methods are used to perform insertion, mutation and iteration operations on an
// unordered collection of Amiibo pointers. Each Amiibo slice can contain 0 to N number of normalized Amiibo, with each
// Amiibo pointer being held in a private slice field. All exposed methods for the Amiibo slice perform a corresponding
// operation for this internal field. This property is protected to prevent incorrect data assignment as the slice permits
// any data interface to be assigned to the Amiibo slice.
type AmiiboSlice struct {
	slice *slice.Slice
}

// Append adds a new Amiibo to the end of Amiibo slice and returns the modified Amiibo slice.
func (pointer *AmiiboSlice) Append(amiibo *Amiibo) *AmiiboSlice {
	pointer.slice.Append(amiibo)
	return pointer
}

// Assign adds N Amiibo to the end Amiibo slice and returns the modified Amiibo slice.
func (pointer *AmiiboSlice) Assign(amiibo ...*Amiibo) *AmiiboSlice {
	for _, amiibo := range amiibo {
		pointer.Append(amiibo)
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the Amiibo slice.
func (pointer *AmiiboSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two Amiibo slices into a single Amiibo slice.
func (pointer *AmiiboSlice) Concatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice {
	pointer.slice.Concatenate(amiiboSlice.slice)
	return pointer
}

// Each executes a provided function once for each element in the Amiibo slice.
func (pointer *AmiiboSlice) Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Amiibo))
	})
	return pointer
}

// EachReverse execute a provided function once for each Amiibo slice in the reverse order found in the Amiibo slice.
func (pointer *AmiiboSlice) EachReverse(f func(i int, amiibo *Amiibo)) *AmiiboSlice {
	pointer.slice.EachReverse(func(i int, value interface{}) {
		f(i, value.(*Amiibo))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Amiibo slice contains zero values.
func (pointer *AmiiboSlice) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the Amiibo pointer held at the argument index. Returns nil if index exceeds Amiibo slice length.
func (pointer *AmiiboSlice) Fetch(i int) *Amiibo {
	amiibo, _ := pointer.Get(i)
	return amiibo
}

// Get returns the Amiibo pointer held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *AmiiboSlice) Get(i int) (*Amiibo, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

// Len method returns the number of elements in the Amiibo slice.
func (pointer *AmiiboSlice) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each Amiibo pointer in the Amiibo slice
// and sets the returned value to the current index.
func (pointer *AmiiboSlice) Map(f func(i int, amiibo *Amiibo) *Amiibo) *AmiiboSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Amiibo))
	})
	return pointer
}

// Poll method removes the first element from the Amiibo slice and returns that removed Amiibo.
func (pointer *AmiiboSlice) Poll() *Amiibo {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Amiibo)
	}
	return nil
}

// Pop method removes the last Amiibo from the Amiibo slice and returns that Amiibo.
func (pointer *AmiiboSlice) Pop() *Amiibo {
	value := pointer.slice.Pop()
	if value != nil {
		return value.(*Amiibo)
	}
	return nil
}

// Preassign method adds zero or more Amiibo pointers to the beginning of the Amiibo slice and returns the modified Amiibo slice.
func (pointer *AmiiboSlice) Preassign(amiibo ...*Amiibo) *AmiiboSlice {
	for _, amiibo := range amiibo {
		pointer.Prepend(amiibo)
	}
	return pointer
}

// Precatenate merges two Amiibo slices, prepending the argument Amiibo slice to the beginning of the receiver Amiibo slice.
func (pointer *AmiiboSlice) Precatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice {
	pointer.slice.Precatenate(amiiboSlice.slice)
	return pointer
}

// Prepend method adds one Amiibo to the beginning of the Amiibo slice and returns the modified Amiibo slice.
func (pointer *AmiiboSlice) Prepend(amiibo *Amiibo) *AmiiboSlice {
	pointer.slice.Prepend(amiibo)
	return pointer
}

// Push method adds a new Amiibo to the end of the Amiibo slice and returns the length of the modified Amiibo slice.
func (pointer *AmiiboSlice) Push(amiibo *Amiibo) int {
	return pointer.slice.Push(amiibo)
}

// Replace method replaces the Amiibo at the argument index if it is in bounds with the provided argument Amiibo.
func (pointer *AmiiboSlice) Replace(i int, amiibo *Amiibo) bool {
	return pointer.slice.Replace(i, amiibo)
}

// Set method returns a unique Amiibo slice, removing duplicate Amiibo that have the same ID.
func (pointer *AmiiboSlice) Set() *AmiiboSlice {
	amiiboSlice := newAmiiboSlice()
	m := map[string]bool{}
	pointer.Each(func(_ int, amiibo *Amiibo) {
		if _, ok := m[amiibo.ID]; !ok {
			m[amiibo.ID] = true
			amiiboSlice.Append(amiibo)
		}
	})
	return amiiboSlice
}

// Slice method returns a shallow copy of a portion of the Amiibo slice into a new Amiibo slice.
// Amiibo slice is selected from begin to end (end not included).
// The original Amiibo slice will not be modified but all values are shared between the two Amiibo slices.
func (pointer *AmiiboSlice) Slice(start, end int) *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.slice.Slice(start, end)}
}

// Splice method changes the contents of the Amiibo slice by removing existing elements fron i to N.
// Returns a new Amiibo slice containing the cut values.
func (pointer *AmiiboSlice) Splice(start, end int) *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.slice.Splice(start, end)}
}

// String returns the string value of the Amiibo slice.
func (pointer *AmiiboSlice) String() string {
	return fmt.Sprintf("%v", *pointer)
}

// Swap swaps the Amiibo pointer held at i to j and vice versa. Does not swap the Amiibo slice
// pointers if either i or j is out of bounds.
func (pointer *AmiiboSlice) Swap(i, j int) *AmiiboSlice {
	pointer.slice.Swap(i, j)
	return pointer
}
