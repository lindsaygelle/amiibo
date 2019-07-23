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

func (pointer *RawAmiiboSlice) Append(rawAmiibo *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Append(rawAmiibo)
	return pointer
}

func (pointer *RawAmiiboSlice) Assign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice {
	for _, rawAmiibo := range rawAmiibo {
		pointer.Append(rawAmiibo)
	}
	return pointer
}

func (pointer *RawAmiiboSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *RawAmiiboSlice) Concatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice {
	pointer.slice.Concatenate(rawAmiiboSlice.slice)
	return pointer
}

func (pointer *RawAmiiboSlice) Each(f func(i int, rawAmiibo *RawAmiibo)) *RawAmiiboSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*RawAmiibo))
	})
	return pointer
}

func (pointer *RawAmiiboSlice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *RawAmiiboSlice) Fetch(i int) *RawAmiibo {
	rawAmiibo, _ := pointer.Get(i)
	return rawAmiibo
}

func (pointer *RawAmiiboSlice) Get(i int) (*RawAmiibo, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*RawAmiibo), ok
	}
	return nil, ok
}

func (pointer *RawAmiiboSlice) Len() int {
	return pointer.slice.Len()
}

func (pointer *RawAmiiboSlice) Map(f func(i int, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*RawAmiibo))
	})
	return pointer
}

func (pointer *RawAmiiboSlice) Poll() *RawAmiibo {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawAmiibo)
	}
	return nil
}

func (pointer *RawAmiiboSlice) Pop() *RawAmiibo {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*RawAmiibo)
	}
	return nil
}

func (pointer *RawAmiiboSlice) Preassign(rawAmiibo ...*RawAmiibo) *RawAmiiboSlice {
	for _, rawAmiibo := range rawAmiibo {
		pointer.Prepend(rawAmiibo)
	}
	return pointer
}

func (pointer *RawAmiiboSlice) Precatenate(rawAmiiboSlice *RawAmiiboSlice) *RawAmiiboSlice {
	pointer.slice.Precatenate(rawAmiiboSlice.slice)
	return pointer
}

func (pointer *RawAmiiboSlice) Prepend(rawAmiibo *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Prepend(rawAmiibo)
	return pointer
}

func (pointer *RawAmiiboSlice) Push(rawAmiibo *RawAmiibo) int {
	return pointer.slice.Push(rawAmiibo)
}

func (pointer *RawAmiiboSlice) Replace(i int, rawAmiibo *RawAmiibo) bool {
	return pointer.slice.Replace(i, rawAmiibo)
}

func (pointer *RawAmiiboSlice) Slice(start, end int) *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *RawAmiiboSlice) Splice(start, end int) *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *RawAmiiboSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
